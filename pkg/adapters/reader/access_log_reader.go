package reader

import (
	"bufio"
	"log"
	"os"
)

// AccessLogReader
type AccessLogReader struct {
}

// NewAccessLogReader
func NewAccessLogReader() *AccessLogReader {
	return &AccessLogReader{}
}

// Read given access log file
func (r *AccessLogReader) Read(filePath string) (<-chan string, <-chan int) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	logChan := make(chan string)
	bytesReadChan := make(chan int)

	go func() {
		defer func() {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			logChan <- scanner.Text()
			bytesReadChan <- len(scanner.Bytes())
		}

		close(logChan)
		close(bytesReadChan)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return logChan, bytesReadChan
}

func (r *AccessLogReader) FileSize(filePath string) int64 {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}

	return fileInfo.Size()
}
