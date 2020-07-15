/*
Author: Costin Canciu
Entrypoint: crypt12-decrypt.go
Description: Decrypt Whatsapp .crypt12 database files
*/

package main

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check_error(err error) {
	if err != nil {
		panic(err)
	}
}

func exit(message string) {
	fmt.Printf(message)
	os.Exit(2)
}

func get_files_arguments() (string, string, string) {
	var keyFile, crypt12File, outputFile string
	flag.StringVar(&keyFile, "keyfile", "key", "decryption key file path")
	flag.StringVar(&crypt12File, "crypt12file", "msgstore.db.crypt12", "crypt12 file path")
	flag.StringVar(&outputFile, "outputfile", "msgstore.db", "decrypted output file path")
	flag.Parse()
	return keyFile, crypt12File, outputFile
}

func fileExists(fileName string) (bool, int64) {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false, 0
	}
	return !info.IsDir(), info.Size()
}

func get_files_size(keyFile, cryptFile string) (int64, int64) {
	keyFileExist, keySize := fileExists(keyFile)
	if !keyFileExist {
		exit("Key file does not exist.")
	}
	if keySize != 158 {
		exit("Key file is invalid. It should be 158 byte size.")
	}
	cryptFileExist, cryptSize := fileExists(cryptFile)
	if !cryptFileExist {
		exit("Crypt12 file does not exist.")
	}
	return keySize, cryptSize
}

func read_key(keyFile string) ([]byte, []byte) {
	fileInput, fileError := os.Open(keyFile)
	defer fileInput.Close()
	check_error(fileError)
	keyBytes := make([]byte, 158)
	_, readError := fileInput.Read(keyBytes)
	check_error(readError)
	t1 := make([]byte, 32)
	copy(t1[:], keyBytes[30:62])
	key := make([]byte, 32)
	copy(key[:], keyBytes[126:158])
	return key, t1
}

func read_crypt12_file(crypt12File string, crypt12Size int64) ([]byte, []byte, []byte) {
	fileInput, fileError := os.Open(crypt12File)
	defer fileInput.Close()
	check_error(fileError)
	crypt12Bytes := make([]byte, crypt12Size)
	_, readError := fileInput.Read(crypt12Bytes)
	check_error(readError)
	t2 := make([]byte, 32)
	copy(t2[:], crypt12Bytes[3:35])
	IV := make([]byte, 16)
	copy(IV[:], crypt12Bytes[51:67])
	cipherText := make([]byte, crypt12Size-67-20)
	copy(cipherText[:], crypt12Bytes[67:crypt12Size-20])
	return cipherText, IV, t2
}

func validate_header(t1, t2 []byte) {
	if string(t1) != string(t2) {
		exit("Key file mismatch or crypt12 file is corrupt.")
	}
}

func crypt12_decrypt(key, IV, cipherText []byte) []byte {
	block, cipherError := aes.NewCipher(key)
	check_error(cipherError)
	aesGCM, gcmError := cipher.NewGCMWithNonceSize(block, 16)
	check_error(gcmError)
	plaintext, encryptError := aesGCM.Open(nil, IV, cipherText, nil)
	check_error(encryptError)
	return plaintext
}

func decompress(compressedBytes []byte) []byte {
	bytesReader := bytes.NewReader(compressedBytes)
	zlibReader, zlibError := zlib.NewReader(bytesReader)
	check_error(zlibError)
	result, inflateError := ioutil.ReadAll(zlibReader)
	check_error(inflateError)
	return result
}

func write_output_file(outputFile string, resultBytes []byte) {
	file, fileError := os.OpenFile(
		outputFile,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	defer file.Close()
	check_error(fileError)
	_, writeError := file.Write(resultBytes)
	check_error(writeError)
}

func validate_sqlite_file(resultBytes []byte) {
	if strings.ToLower(string(resultBytes[:6])) != "sqlite" {
		exit("Decryption failed. SQLite file format is corrupt.")
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Decryption failed. Error: %v \n", r)
		}
	}()

	keyFile, crypt12File, outputFile := get_files_arguments()

	_, crypt12Size := get_files_size(keyFile, crypt12File)

	key, t1 := read_key(keyFile)

	cipherText, IV, t2 := read_crypt12_file(crypt12File, crypt12Size)

	validate_header(t1, t2)

	plainText := crypt12_decrypt(key, IV, cipherText)

	resultBytes := decompress(plainText)

	validate_sqlite_file(resultBytes)

	write_output_file(outputFile, resultBytes)

	fmt.Printf("Decryption successful: %s", outputFile)
}
