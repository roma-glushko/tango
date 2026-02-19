package component

import (
	"fmt"
	"log"
	"os"
	"tango/pkg/services"

	"github.com/cheggaaa/pb"
)

// ReaderProgressDecorator decorates an AccessLogReader with a progress bar.
type ReaderProgressDecorator struct {
	accessLogReader services.AccessLogReader
	progressBar     *pb.ProgressBar
}

// NewReaderProgressDecorator creates a new ReaderProgressDecorator instance.
func NewReaderProgressDecorator(accessLogReader services.AccessLogReader) *ReaderProgressDecorator {
	return &ReaderProgressDecorator{
		accessLogReader: accessLogReader,
		progressBar:     nil,
	}
}

// Read opens a file, displays a progress bar, and delegates reading to the wrapped reader.
func (r *ReaderProgressDecorator) Read(filePath string, readAccessLogFunc services.ReadAccessLogFunc) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = file.Close() }()

	fileInfo, err := file.Stat()

	if err != nil {
		log.Println(err)
		return
	}

	fileSize := fileInfo.Size()

	r.progressBar = pb.New64(fileSize)

	r.progressBar.Format("[=>_]")
	r.progressBar.SetUnits(pb.U_BYTES)
	r.progressBar.SetMaxWidth(100)

	r.start()

	r.accessLogReader.Read(filePath, func(line []byte, byteCount int) {
		readAccessLogFunc(line, byteCount)

		r.update(byteCount)
	})

	r.progressBar.Finish()
}

// Map delegates to the underlying reader if it supports memory mapping.
func (r *ReaderProgressDecorator) Map(filePath string) ([]byte, func(), error) {
	if mr, ok := r.accessLogReader.(services.MappableReader); ok {
		return mr.Map(filePath)
	}
	return nil, nil, fmt.Errorf("underlying reader does not support mapping")
}

func (r *ReaderProgressDecorator) start() {
	r.progressBar.Start()
}

func (r *ReaderProgressDecorator) update(byteCount int) {
	r.progressBar.Add(byteCount)
}
