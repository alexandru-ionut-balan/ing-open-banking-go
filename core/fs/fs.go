package fs

import (
	"bufio"
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

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Warn.Println("Cannot close file "+file.Name(), err)
	}
}
