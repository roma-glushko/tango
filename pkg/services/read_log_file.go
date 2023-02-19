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
	logMapper        *mapper.AccessLogMapper
	logReader        *adapters.AccessLogReader
	readProgress     *adapters.ReadProgress
	filterLogService FilterAccessLogService
	ipProcessor      processor.IPProcessor
}

// NewReadAccessLogService Create a new ReadAccessLogService
func NewReadAccessLogService(
	logMapper *mapper.AccessLogMapper,
	accessLogReader *adapters.AccessLogReader,
	readProgress *adapters.ReadProgress,
	filterLogService FilterAccessLogService,
	ipProcessor processor.IPProcessor,
) *ReadAccessLogService {
	return &ReadAccessLogService{
		logMapper:        logMapper,
		logReader:        accessLogReader,
		readProgress:     readProgress,
		filterLogService: filterLogService,
		ipProcessor:      ipProcessor,
	}
}

// Read Access Logs (filter them out if needed) and add to the channel
func (s *ReadAccessLogService) Read(filePath string) <-chan entity.AccessLogRecord {
	logFileSize := s.logReader.FileSize(filePath)
	rawLogChan, bytesReadChan := s.logReader.Read(filePath)

	parsedLogChan := make(chan entity.AccessLogRecord)

	s.readProgress.Track(bytesReadChan, logFileSize)

	var waitGroup sync.WaitGroup

	// TODO: Configure thread number
	for i := 0; i < 4; i++ {
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			for rawLog := range rawLogChan {
				logRecord := s.logMapper.Map(rawLog)

				// process a parsed access log record
				logRecord = s.ipProcessor.Process(logRecord)

				// filter/skip parsed access log record if needed
				if s.filterLogService.Filter(logRecord) {
					s.logMapper.Recycle(logRecord)

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
