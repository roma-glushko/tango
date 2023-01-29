package reader

import (
	"github.com/cheggaaa/pb"
)

type ReadProgress struct {
	progressBar *pb.ProgressBar
}

//
func NewReadProgress() *ReadProgress {
	return &ReadProgress{
		progressBar: nil,
	}
}

func (r *ReadProgress) Track(bytesReadChan <-chan int, fileSize int64) {
	r.progressBar = pb.New64(fileSize)

	r.progressBar.Format("[=>_]")
	r.progressBar.SetUnits(pb.U_BYTES)
	r.progressBar.SetMaxWidth(100)

	go func() {
		r.progressBar.Start()
		defer r.progressBar.Finish()

		for bytesRead := range bytesReadChan {
			r.progressBar.Add(bytesRead)
		}
	}()
}
