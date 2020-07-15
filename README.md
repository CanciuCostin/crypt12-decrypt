# Crypt12 Decrypt

![GO][go-shield]

Decrypt Whatsapp crypt12  sqlite database files

## Crypt12 algorithm

* AES GCM mode encryption using 128 bit block size and 16 bytes IV (nonce)

* key file must be 158 byte long
![Key](key.png?raw=true "Key")

* crypt12 file includes 57 byte header and 20 byte trailer which needs to be removed
![crypt12file](crypt12file.png?raw=true "crypt12 file")

* the resulted plain bytes need to be decompressed in order to obtain the final SQLite .db file

## Usage example

1. Ensure you have the key file (**key**) and crypt12 file (**msgstore.db.crypt12**) in the same directory with the go entrypoint
* Run using GO:
```
go run crypt12-decrypt
```
* Or run Windows executable:
```
crypt12-decrypt.exe
```
 2. Otherwise use the necessery arguments:
```
go run crypt12-decrypt.go -h
```
```
Usage of crypt12-decrypt.exe:
  -crypt12file string
        crypt12 file path (default "msgstore.db.crypt12")
  -keyfile string
        decryption key file path (default "key")
  -outputfile string
        decrypted output file path (default "msgstore.db")
```

## Build

```
go get github.com/CanciuCostin/crypt12-decrypt

go build crypt12-decrypt -> built with go version 1.14.4
```


<!-- Markdown link & img dfn's -->
[go-shield]: https://img.shields.io/badge/go-1.14.4-green
