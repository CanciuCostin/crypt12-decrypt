![GO][go-shield]
# Crypt12 Decrypt

Decrypt Whatsapp crypt12  sqlite database files

## Crypt12 algorithm

* AES GCM mode encryption using 128 bit block size and 16 bytes IV (nonce)

* key file must be 158 byte long (only last 32bytes represent the key)

![Key](key.png?raw=true "Key")

* crypt12 file includes 57 byte header and 20 byte trailer which needs to be removed

![crypt12file](crypt12file.png?raw=true "crypt12 file")

* the resulted plain bytes need to be decompressed in order to obtain the final SQLite .db file

* for more information you can check [A Systems Approach to Cyber Security: Proceedings of the 2nd Singapore Cyber-Security R&D Conference (SG-CRC 2017)](https://books.google.ro/books?id=RUXiDgAAQBAJ&pg=PR7&lpg=PR7&dq=A+Systems+Approach+to+Cyber+Security:+Proceedings+of+the+2nd+Singapore+Cyber-Security+R%26D+Conference+(SG-CRC+2017)&source=bl&ots=vWJcT_nFMa&sig=ACfU3U1Tmj6ui8lYaPPZIY7fGz4UwAIN6w&hl=en&sa=X&ved=2ahUKEwjKpannwc_qAhWok4sKHTrNDtAQ6AEwBXoECBIQAQ#v=onepage&q=A%20Systems%20Approach%20to%20Cyber%20Security%3A%20Proceedings%20of%20the%202nd%20Singapore%20Cyber-Security%20R%26D%20Conference%20(SG-CRC%202017)&f=false)

## Usage example

1. Ensure you have the key file (**key**) and crypt12 file (**msgstore.db.crypt12**) in the same directory with the go entrypoint. You can use the existing key and msgstore files in the repo for testing.
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
