package di

import (
	"tango/internal/cli/component"
	"tango/internal/infrastructure/reader"
	"tango/internal/usecase"
)

//
func InitReadAccessLogUsecase() *usecase.ReadAccessLogUsecase {
	accessLogReader := reader.NewAccessLogReader()
	readerProgressDecorator := component.NewReaderProgressDecorator(accessLogReader)

	return usecase.NewReadAccessLogUsecase(readerProgressDecorator)
}
