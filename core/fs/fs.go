package fs

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"strings"

	"github.com/alexandru-ionut-balan/ing-open-banking-go/core/log"
)

func ReadFileText(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Error.Println("Cannot open file at"+filePath+" for reading", err)
		return ""
	}

	defer closeFile(file)

	scanner := bufio.NewScanner(file)
	text := strings.Builder{}

	for scanner.Scan() {
		text.WriteString(scanner.Text())
	}

	return text.String()
}

func ReadFileBytes(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		log.Error.Println("Cannot open file at"+filePath+" for reading", err)
		return nil
	}

	defer closeFile(file)

	stat, err := file.Stat()
	if err != nil {
		log.Error.Println("Cannot stat file "+file.Name(), err)
		return nil
	}

	bytes := make([]byte, stat.Size())

	_, err = file.Read(bytes)
	if err != nil {
		log.Error.Println("Cannot read file "+file.Name(), err)
		return nil
	}

	return bytes
}

func ReadCertificate(filePath string) *x509.Certificate {
	bytes := ReadFileBytes(filePath)

	derBytes, rest := pem.Decode(bytes)
	if len(rest) > 0 {
		log.Error.Fatalln("Cannot read certificate at " + filePath)
	}

	certificate, err := x509.ParseCertificate(derBytes.Bytes)
	if err != nil {
		log.Error.Fatalln("Cannot parse certificate from file at "+filePath, err)
	}

	return certificate
}

func ReadPrivateKey(filePath string) *rsa.PrivateKey {
	bytes := ReadFileBytes(filePath)

	derBytes, rest := pem.Decode(bytes)
	if len(rest) > 0 {
		log.Error.Fatalln("Cannot read certificate at " + filePath)
	}

	privateKey, err := readPkcs1(derBytes.Bytes)
	if err != nil {
		log.Error.Println(err)
		privateKey, err = readPkcs8(derBytes.Bytes)
		if err != nil {
			log.Error.Fatalln(err)
		}

		return privateKey
	}

	return privateKey
}

func readPkcs1(derBytes []byte) (*rsa.PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey(derBytes)
}

func readPkcs8(derBytes []byte) (*rsa.PrivateKey, error) {
	privateKey, err := x509.ParsePKCS8PrivateKey(derBytes)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("cannot convert the read private key to the RSA format")
	}

	return rsaKey, nil
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Warn.Println("Cannot close file "+file.Name(), err)
	}
}
