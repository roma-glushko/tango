package component

import (
	"log"
	"os"
	"tango/internal/infrastructure/reader"
	"tango/internal/usecase"

	"github.com/cheggaaa/pb"
)

type ReaderProgressDecorator struct {
	accessLogReader reader.AccessLogReader
	progressBar     *pb.ProgressBar
}

//
func NewReaderProgressDecorator(accessLogReader *reader.AccessLogReader) *ReaderProgressDecorator {
	return &ReaderProgressDecorator{
		accessLogReader: *accessLogReader,
		progressBar:     nil,
	}
}

func (r *ReaderProgressDecorator) Read(filePath string, readAccessLogFunc usecase.ReadAccessLogFunc) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}

	fileSize := fileInfo.Size()

	r.progressBar = pb.New64(fileSize)

	r.start()

	r.accessLogReader.Read(filePath, func(accessLogRecord string, byteCount int) {
		readAccessLogFunc(accessLogRecord, byteCount)

		r.update(byteCount)
	})

	r.progressBar.Finish()
}

//
func (r *ReaderProgressDecorator) start() {
	r.progressBar.Start()
}

//
func (r *ReaderProgressDecorator) update(byteCount int) {
	r.progressBar.Add(byteCount)
}
