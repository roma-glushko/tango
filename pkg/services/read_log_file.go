package services

import (
	"sync"
	adapters "tango/pkg/adapters/reader"
	"tango/pkg/entity"
	"tango/pkg/services/mapper"
	"tango/pkg/services/processor"
)

//
type ReadAccessLogFunc func(accessLogRecord string, bytes int)

//
type ReadAccessLogService struct {
	logReader        *adapters.AccessLogReader
	readProgress     *adapters.ReadProgress
	filterLogService FilterAccessLogService
	ipProcessor      processor.IPProcessor
}

// NewReadAccessLogService Create a new ReadAccessLogService
func NewReadAccessLogService(
	accessLogReader *adapters.AccessLogReader,
	readProgress *adapters.ReadProgress,
	filterLogService FilterAccessLogService,
	ipProcessor processor.IPProcessor,
) *ReadAccessLogService {
	return &ReadAccessLogService{
		logReader:        accessLogReader,
		readProgress:     readProgress,
		filterLogService: filterLogService,
		ipProcessor:      ipProcessor,
	}
}

// Read Access Logs and convert them to array of AccessLogRecord-s
func (u *ReadAccessLogService) Read(filePath string) <-chan entity.AccessLogRecord {
	logFileSize := u.logReader.FileSize(filePath)
	rawLogChan, bytesReadChan := u.logReader.Read(filePath)

	parsedLogChan := make(chan entity.AccessLogRecord)

	u.readProgress.Track(bytesReadChan, logFileSize)

	var waitGroup sync.WaitGroup

	// TODO: Configure thread number
	for i := 0; i < 4; i++ {
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			for rawLog := range rawLogChan {
				logRecord := mapper.MapAccessLogRecord(rawLog)

				// process a parsed access log record
				logRecord = u.ipProcessor.Process(logRecord)

				// filter/skip parsed access log record if needed
				if u.filterLogService.Filter(logRecord) {
					continue
				}

				parsedLogChan <- logRecord
			}
		}()
	}

	go func() {
		waitGroup.Wait()
		close(parsedLogChan)
	}()

	return parsedLogChan
}
