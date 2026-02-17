package component

import (
	"log"
	"os"
	"tango/pkg/adapters/reader"
	"tango/pkg/services"

	"github.com/cheggaaa/pb"
)

// ReaderProgressDecorator decorates an AccessLogReader with a progress bar.
type ReaderProgressDecorator struct {
	accessLogReader reader.AccessLogReader
	progressBar     *pb.ProgressBar
}

// NewReaderProgressDecorator creates a new ReaderProgressDecorator instance.
func NewReaderProgressDecorator(accessLogReader *reader.AccessLogReader) *ReaderProgressDecorator {
	return &ReaderProgressDecorator{
		accessLogReader: *accessLogReader,
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

	r.accessLogReader.Read(filePath, func(accessLogRecord string, byteCount int) {
		readAccessLogFunc(accessLogRecord, byteCount)

		r.update(byteCount)
	})

	r.progressBar.Finish()
}

func (r *ReaderProgressDecorator) start() {
	r.progressBar.Start()
}

func (r *ReaderProgressDecorator) update(byteCount int) {
	r.progressBar.Add(byteCount)
}
