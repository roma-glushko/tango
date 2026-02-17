package writer

import (
	"bufio"
	"encoding/csv"
	"os"
)

// newBufferedCsvWriter creates a csv.Writer backed by a buffered writer of the given size.
// The caller must call Flush() on the returned csv.Writer and
// Flush() on the returned bufio.Writer before closing the file.
func newBufferedCsvWriter(file *os.File, bufferSize int) (*csv.Writer, *bufio.Writer) {
	buffered := bufio.NewWriterSize(file, bufferSize)
	return csv.NewWriter(buffered), buffered
}
