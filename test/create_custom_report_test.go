package test

import (
	"tango/internal/cli"
	"tango/internal/infrastructure/writer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCustomReportWithKeepIPFilter(t *testing.T) {
	assert := assert.New(t)
	tangoCli := cli.NewTangoCli()

	sampleIP := "130.93.253.236"
	reportFilePath := "results/custom-report-keep-ip-filter.csv"

	tangoCli.Run([]string{
		"main",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
		"-c",
		"fixture/.tango.empty.yaml",
		"--keep-ip-filter",
		sampleIP,
		"custom",
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	reportHeader, reportBody := testReport[0], testReport[1:]

	assert.Equal(reportHeader, writer.CustomReportHeader, "Custom Report Header is not the same")
	assert.Equal(18, len(reportBody), "Report File should contain 18 records")

	for _, reportItem := range reportBody {
		assert.Contains(reportItem[1], sampleIP, "Report Items should be from filtered IP")
	}
}

func TestCreateCustomReportWithKeepUAFilter(t *testing.T) {
	assert := assert.New(t)
	tangoCli := cli.NewTangoCli()

	sampleUserAgent := "iPhone OS 12_3_1 like Mac OS X"
	reportFilePath := "results/custom-report-keep-ua-filter.csv"

	tangoCli.Run([]string{
		"main",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
		"-c",
		"fixture/.tango.empty.yaml",
		"--keep-ua-filter",
		sampleUserAgent,
		"custom",
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	t.Logf("%v", len(testReport))

	reportHeader, reportBody := testReport[0], testReport[1:]

	assert.Equal(reportHeader, writer.CustomReportHeader, "Custom Report Header is not the same")
	assert.Equal(len(reportBody), 50, "Report File should contain 50 records")

	for _, reportItem := range reportBody {
		assert.Contains(reportItem[5], sampleUserAgent, "Report Items should be created from filtered UA")
	}
}
