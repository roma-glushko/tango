package writer

import (
	"bufio"
	"encoding/csv"
	"os"
)

const csvWriteBufferSize = 256 * 1024 // 256KB

// newBufferedCsvWriter creates a csv.Writer backed by a 256KB buffered writer.
// The caller must call Flush() on the returned csv.Writer and
// Flush() on the returned bufio.Writer before closing the file.
func newBufferedCsvWriter(file *os.File) (*csv.Writer, *bufio.Writer) {
	buffered := bufio.NewWriterSize(file, csvWriteBufferSize)
	return csv.NewWriter(buffered), buffered
}
