package core

import (
	"bufio"
	"os"
	"strings"

	"github.com/alexandru-ionut-balan/ing-open-banking-go/logger"
)

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		logger.Error("Unable to close file "+file.Name(), err)
	}
}

// ReadFileAsString tries to read the file at {path}. If it cannot read, a non nil
// error will be returned and the output string will be the empty string.
func ReadFileAsString(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer closeFile(file)

	scanner := bufio.NewScanner(file)
	fileText := strings.Builder{}

	for scanner.Scan() {
		fileText.WriteString(scanner.Text())
	}

	return fileText.String(), nil
}

// ReadFileAsBytes tries to read the file at {path}. If it cannot read, a non nil
// error will be returned and the output byte array will be empty.
func ReadFileAsBytes(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer closeFile(file)

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	bytes := make([]byte, stat.Size())
	_, err = file.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
