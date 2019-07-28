package reader

import (
	"bufio"
	"log"
	"os"
	"tango/internal/usecase"
)

type AccessLogReader struct {
}

//
func NewAccessLogReader() *AccessLogReader {
	return &AccessLogReader{}
}

// Read given access log file
func (r *AccessLogReader) Read(filePath string, readAccessLogFunc usecase.ReadAccessLogFunc) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		readAccessLogFunc(scanner.Text(), len(scanner.Bytes()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
