package reader

import (
	"bytes"
	"log"
	"os"
	"tango/pkg/services"

	"github.com/edsrzf/mmap-go"
)

// MmapAccessLogReader reads access log files using memory-mapped I/O.
type MmapAccessLogReader struct{}

// NewMmapAccessLogReader creates a new MmapAccessLogReader instance.
func NewMmapAccessLogReader() *MmapAccessLogReader {
	return &MmapAccessLogReader{}
}

// Map memory-maps the file and returns raw data with a cleanup function.
func (r *MmapAccessLogReader) Map(filePath string) ([]byte, func(), error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	data, err := mmap.Map(f, mmap.RDONLY, 0)
	if err != nil {
		if cerr := f.Close(); cerr != nil {
			log.Println("Error closing file:", cerr)
		}
		return nil, nil, err
	}

	cleanup := func() {
		if err := data.Unmap(); err != nil {
			log.Println("Error unmapping file:", err)
		}
		if err := f.Close(); err != nil {
			log.Println("Error closing file:", err)
		}
	}

	return data, cleanup, nil
}

// Read memory-maps the file and calls readAccessLogFunc for each line.
func (r *MmapAccessLogReader) Read(filePath string, readAccessLogFunc services.ReadAccessLogFunc) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data, err := mmap.Map(f, mmap.RDONLY, 0)
	if err != nil {
		if cerr := f.Close(); cerr != nil {
			log.Println("Error closing file:", cerr)
		}
		log.Fatal(err)
	}
	defer func() {
		if err := data.Unmap(); err != nil {
			log.Println("Error unmapping file:", err)
		}
		if err := f.Close(); err != nil {
			log.Println("Error closing file:", err)
		}
	}()

	offset := 0
	for offset < len(data) {
		nl := bytes.IndexByte(data[offset:], '\n')
		var line []byte
		if nl < 0 {
			line = data[offset:]
			offset = len(data)
		} else {
			line = data[offset : offset+nl]
			offset += nl + 1
		}
		if len(line) > 0 && line[len(line)-1] == '\r' {
			line = line[:len(line)-1]
		}
		if len(line) == 0 {
			continue
		}
		readAccessLogFunc(line, len(line)+1)
	}
}
