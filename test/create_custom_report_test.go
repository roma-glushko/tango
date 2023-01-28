package test

import (
	"net"
	"strings"
	"tango/pkg/cli"
	"tango/pkg/infrastructure/writer"
	"tango/pkg/services/filter"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateCustomReportWithoutFilters(t *testing.T) {
	assert := assert.New(t)
	tangoCli := cli.NewTangoCli("0.0.0-test", "dummycommithash")

	reportFilePath := "results/custom-report-keep-ip-filter.csv"

	tangoCli.Run([]string{
		"main",
		"-c",
		"fixture/.tango.empty.yaml",
		"custom",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	reportHeader, reportBody := testReport[0], testReport[1:]

	assert.Equal(reportHeader, writer.CustomReportHeader, "Custom Report Header is not the same")
	assert.Len(reportBody, 200, "Report file should be as full as original input file")
}

func TestCreateCustomReportWithKeepIPFilter(t *testing.T) {
	assert := assert.New(t)
	tangoCli := cli.NewTangoCli("0.0.0-test", "dummycommithash")

	sampleIP := "130.93.253.236"
	reportFilePath := "results/custom-report-keep-ip-filter.csv"

	tangoCli.Run([]string{
		"main",
		"-c",
		"fixture/.tango.empty.yaml",
		"--keep-ip-filter",
		sampleIP,
		"custom",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
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
	tangoCli := cli.NewTangoCli("0.0.0-test", "dummycommithash")

	sampleUserAgent := "iPhone OS 12_3_1 like Mac OS X"
	reportFilePath := "results/custom-report-keep-ua-filter.csv"

	tangoCli.Run([]string{
		"main",
		"-c",
		"fixture/.tango.empty.yaml",
		"--keep-ua-filter",
		sampleUserAgent,
		"custom",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	reportHeader, reportBody := testReport[0], testReport[1:]

	assert.Equal(reportHeader, writer.CustomReportHeader, "Custom Report Header is not the same")
	assert.Equal(len(reportBody), 50, "Report File should contain 50 records")

	for _, reportItem := range reportBody {
		assert.Contains(reportItem[5], sampleUserAgent, "Report Item should contain needed User Agent")
	}
}

func TestCreateCustomReportWithUAFilter(t *testing.T) {
	assert := assert.New(t)
	tangoCli := cli.NewTangoCli("0.0.0-test", "dummycommithash")

	sampleUserAgent := "iPhone OS 12_3_1 like Mac OS X"
	reportFilePath := "results/custom-report-ua-filter.csv"

	tangoCli.Run([]string{
		"main",
		"--ua-filter",
		sampleUserAgent,
		"custom",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	reportHeader, reportBody := testReport[0], testReport[1:]

	assert.Equal(reportHeader, writer.CustomReportHeader, "Custom Report Header is not the same")

	for _, reportItem := range reportBody {
		assert.NotContains(reportItem[5], sampleUserAgent, "Needed user agent should be filtered")
	}
}

func TestCreateCustomReportWithKeepUriFilter(t *testing.T) {
	assert := assert.New(t)
	tangoCli := cli.NewTangoCli("0.0.0-test", "dummycommithash")

	sampleURI := "/category200"
	reportFilePath := "results/custom-report-keep-uri-filter.csv"

	tangoCli.Run([]string{
		"main",
		"--keep-uri-filter",
		sampleURI,
		"custom",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	reportHeader, reportBody := testReport[0], testReport[1:]

	assert.Equal(reportHeader, writer.CustomReportHeader, "Custom Report Header is not the same")
	assert.Equal(len(reportBody), 2, "Report File should contain 2 records")

	for _, reportItem := range reportBody {
		assert.Contains(reportItem[2], sampleURI, "Report Item should contain required URI")
	}
}

func TestCreateCustomReportWithUriFilter(t *testing.T) {
	assert := assert.New(t)
	tangoCli := cli.NewTangoCli("0.0.0-test", "dummycommithash")

	sampleURI := "/category200"
	reportFilePath := "results/custom-report-uri-filter.csv"

	tangoCli.Run([]string{
		"main",
		"--uri-filter",
		sampleURI,
		"custom",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	reportHeader, reportBody := testReport[0], testReport[1:]

	assert.Equal(reportHeader, writer.CustomReportHeader, "Custom Report Header is not the same")

	for _, reportItem := range reportBody {
		assert.NotContains(reportItem[2], sampleURI, "Needed URI should be filtered")
	}
}

func TestCreateCustomReportWithMultipleAssetFilters(t *testing.T) {
	assert := assert.New(t)
	tangoCli := cli.NewTangoCli("0.0.0-test", "dummycommithash")

	assetPattern1 := "/pub/static/"
	assetPattern2 := "/pub/media/"

	reportFilePath := "results/custom-report-with-multiple-asset-filters.csv"

	tangoCli.Run([]string{
		"main",
		"--asset-filter",
		assetPattern1,
		"--asset-filter",
		assetPattern2,
		"custom",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	reportHeader, reportBody := testReport[0], testReport[1:]

	assert.Equal(reportHeader, writer.CustomReportHeader, "Custom Report Header is not the same")

	for _, reportItem := range reportBody {
		assert.NotContains(reportItem[2], assetPattern1, "Assets should be filtered")
		assert.NotContains(reportItem[2], assetPattern2, "Assets should be filtered")
	}
}

func TestCreateCustomReportWithKeepTimeFilter(t *testing.T) {
	assert := assert.New(t)
	tangoCli := cli.NewTangoCli("0.0.0-test", "dummycommithash")

	testPeriodStart := "2019-07-08 00:00:00 -0200"
	testPeriodFrameEnd := "2019-07-08 00:00:20 -0200"

	reportFilePath := "results/custom-report-with-keep-time-filter.csv"

	tangoCli.Run([]string{
		"main",
		"--keep-time-filter",
		testPeriodStart,
		"--keep-time-filter",
		testPeriodFrameEnd,
		"custom",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
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
	tangoCli := cli.NewTangoCli("0.0.0-test", "dummycommithash")

	reportFilePath := "results/custom-report-with-system-ips-processor.csv"
	systemIPSubnet1 := "157.52.64.0/18"
	systemIP2 := "104.156.90.48"

	tangoCli.Run([]string{
		"main",
		"--system-ips",
		systemIPSubnet1,
		"--system-ips",
		systemIP2,
		"custom",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
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
	tangoCli := cli.NewTangoCli("0.0.0-test", "dummycommithash")

	reportFilePath := "results/custom-report-with-system-ips-processor.csv"
	systemIPSubnet := "157.52.64.0/18"

	tangoCli.Run([]string{
		"main",
		"--system-ips",
		systemIPSubnet,
		"custom",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
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
	tangoCli := cli.NewTangoCli("0.0.0-test", "dummycommithash")

	reportFilePath := "results/custom-report-with-system-ips-processor.csv"
	systemIP := "104.156.90.48"

	tangoCli.Run([]string{
		"main",
		"--system-ips",
		systemIP,
		"custom",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
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
