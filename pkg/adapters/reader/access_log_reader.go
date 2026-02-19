package reader

import (
	"bufio"
	"log"
	"os"
	"tango/pkg/services"
)

// AccessLogReader reads access log files line by line.
type AccessLogReader struct {
	bufferSize int
}

// NewAccessLogReader creates a new AccessLogReader instance with a configurable buffer size.
func NewAccessLogReader(bufferSize int) *AccessLogReader {
	return &AccessLogReader{bufferSize: bufferSize}
}

// Read opens and reads a given access log file, calling readAccessLogFunc for each line.
func (r *AccessLogReader) Read(filePath string, readAccessLogFunc services.ReadAccessLogFunc) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)

	if r.bufferSize > 0 {
		buf := make([]byte, r.bufferSize)
		scanner.Buffer(buf, r.bufferSize)
	}

	for scanner.Scan() {
		line := scanner.Bytes()
		lineCopy := make([]byte, len(line))
		copy(lineCopy, line)
		readAccessLogFunc(lineCopy, len(line))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		return
	}
}
