package test

import (
	"net"
	"strings"
	"tango/internal/cli"
	"tango/internal/infrastructure/writer"
	"tango/internal/usecase/filter"
	"testing"
	"time"

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
		assert.Contains(reportItem[1], sampleIP, "Report Item should contain needed IP")
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
		assert.Contains(reportItem[5], sampleUserAgent, "Report Item should contain needed User Agent")
	}
}

func TestCreateCustomReportWithKeepUriFilter(t *testing.T) {
	assert := assert.New(t)
	tangoCli := cli.NewTangoCli()

	sampleURI := "/category200"
	reportFilePath := "results/custom-report-keep-uri-filter.csv"

	tangoCli.Run([]string{
		"main",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
		"-c",
		"fixture/.tango.empty.yaml",
		"--keep-uri-filter",
		sampleURI,
		"custom",
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	t.Logf("%v", len(testReport))

	reportHeader, reportBody := testReport[0], testReport[1:]

	assert.Equal(reportHeader, writer.CustomReportHeader, "Custom Report Header is not the same")
	assert.Equal(len(reportBody), 2, "Report File should contain 2 records")

	for _, reportItem := range reportBody {
		assert.Contains(reportItem[2], sampleURI, "Report Item should contain required URI")
	}
}

func TestCreateCustomReportWithMultipleAssetFilters(t *testing.T) {
	assert := assert.New(t)
	tangoCli := cli.NewTangoCli()

	assetPattern1 := "/pub/static/"
	assetPattern2 := "/pub/media/"

	reportFilePath := "results/custom-report-with-multiple-asset-filters.csv"

	tangoCli.Run([]string{
		"main",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
		"-c",
		"fixture/.tango.empty.yaml",
		"--asset-filter",
		assetPattern1,
		"--asset-filter",
		assetPattern2,
		"custom",
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	t.Logf("%v", len(testReport))

	reportHeader, reportBody := testReport[0], testReport[1:]

	assert.Equal(reportHeader, writer.CustomReportHeader, "Custom Report Header is not the same")

	for _, reportItem := range reportBody {
		assert.NotContains(reportItem[2], assetPattern1, "Report Item should not be an asset")
		assert.NotContains(reportItem[2], assetPattern2, "Report Item should not be an asset")
	}
}

func TestCreateCustomReportWithKeepTimeFilter(t *testing.T) {
	assert := assert.New(t)
	tangoCli := cli.NewTangoCli()

	testPeriodStart := "2019-07-08 00:00:00 -0200"
	testPeriodFrameEnd := "2019-07-08 00:00:20 -0200"

	reportFilePath := "results/custom-report-with-keep-time-filter.csv"

	tangoCli.Run([]string{
		"main",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
		"-c",
		"fixture/.tango.empty.yaml",
		"--keep-time-filter",
		testPeriodStart,
		"--keep-time-filter",
		testPeriodFrameEnd,
		"custom",
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	reportHeader, reportBody := testReport[0], testReport[1:]

	assert.Equal(reportHeader, writer.CustomReportHeader, "Custom Report Header is not the same")

	timeFrameStart, _ := time.Parse(filter.EuropeFormat, testPeriodStart)
	timeFrameEnd, _ := time.Parse(filter.EuropeFormat, testPeriodFrameEnd)

	for _, reportItem := range reportBody {
		reportItemTime, _ := time.Parse(filter.EuropeFormat, reportItem[0]) // todo: do we need this exactly format?

		assert.True(reportItemTime.After(timeFrameStart) && reportItemTime.Before(timeFrameEnd), "Time of Report Item should within the test duration")
	}
}

func TestCreateCustomReportWithMultipleSystemIpProcessor(t *testing.T) {
	assert := assert.New(t)
	tangoCli := cli.NewTangoCli()

	reportFilePath := "results/custom-report-with-system-ips-processor.csv"
	systemIPSubnet1 := "157.52.64.0/18"
	systemIP2 := "104.156.90.48"

	tangoCli.Run([]string{
		"main",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
		"-c",
		"fixture/.tango.empty.yaml",
		"--system-ips",
		systemIPSubnet1,
		"--system-ips",
		systemIP2,
		"custom",
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	reportHeader, reportBody := testReport[0], testReport[1:]

	assert.Equal(reportHeader, writer.CustomReportHeader, "Custom Report Header is not the same")

	_, IPSubnet1, _ := net.ParseCIDR(systemIPSubnet1)

	// Check if IP patterns work in the output report
	for _, reportItem := range reportBody {
		ipList := strings.Split(reportItem[1], ", ")

		for _, ip := range ipList {
			parsedIP := net.ParseIP(ip)

			assert.NotEqual(systemIP2, ip, "Single IP pattern should filter all related IPs")
			assert.False(IPSubnet1.Contains(parsedIP), "Subnet IP pattern should filter all related IPs")
		}
	}
}

func TestCreateCustomReportWithSubnetSystemIpProcessor(t *testing.T) {
	assert := assert.New(t)
	tangoCli := cli.NewTangoCli()

	reportFilePath := "results/custom-report-with-system-ips-processor.csv"
	systemIPSubnet := "157.52.64.0/18"

	tangoCli.Run([]string{
		"main",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
		"-c",
		"fixture/.tango.empty.yaml",
		"--system-ips",
		systemIPSubnet,
		"custom",
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	reportHeader, reportBody := testReport[0], testReport[1:]

	assert.Equal(reportHeader, writer.CustomReportHeader, "Custom Report Header is not the same")

	_, IPSubnet, _ := net.ParseCIDR(systemIPSubnet)

	// Check if IP pattern works in the output report
	for _, reportItem := range reportBody {
		ipList := strings.Split(reportItem[1], ", ")

		for _, ip := range ipList {
			parsedIP := net.ParseIP(ip)

			assert.False(IPSubnet.Contains(parsedIP), "Subnet IP pattern should filter all related IPs")
		}
	}
}

func TestCreateCustomReportWithSingleSystemIpProcessor(t *testing.T) {
	assert := assert.New(t)
	tangoCli := cli.NewTangoCli()

	reportFilePath := "results/custom-report-with-system-ips-processor.csv"
	systemIP := "104.156.90.48"

	tangoCli.Run([]string{
		"main",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
		"-c",
		"fixture/.tango.empty.yaml",
		"--system-ips",
		systemIP,
		"custom",
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	reportHeader, reportBody := testReport[0], testReport[1:]

	assert.Equal(reportHeader, writer.CustomReportHeader, "Custom Report Header is not the same")

	// Check if IP pattern works in the output report
	for _, reportItem := range reportBody {
		ipList := strings.Split(reportItem[1], ", ")

		for _, ip := range ipList {
			assert.NotEqual(systemIP, ip, "Single IP pattern should filter all related IPs")
		}
	}
}
