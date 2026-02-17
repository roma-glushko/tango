package reader

import (
	"bufio"
	"log"
	"os"
	"tango/pkg/services"
)

// AccessLogReader reads access log files line by line.
type AccessLogReader struct {
}

// NewAccessLogReader creates a new AccessLogReader instance.
func NewAccessLogReader() *AccessLogReader {
	return &AccessLogReader{}
}

// Read opens and reads a given access log file, calling readAccessLogFunc for each line.
func (r *AccessLogReader) Read(filePath string, readAccessLogFunc services.ReadAccessLogFunc) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		readAccessLogFunc(scanner.Text(), len(scanner.Bytes()))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		return
	}
}
