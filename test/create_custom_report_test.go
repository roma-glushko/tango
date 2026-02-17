package test

import (
	"net"
	"strings"
	"tango/pkg/adapters/writer"
	"tango/pkg/cli"
	"tango/pkg/services/filter"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateCustomReportWithoutFilters(t *testing.T) {
	require := require.New(t)
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

	require.Equal(reportHeader, writer.CustomReportHeader)
	require.Len(reportBody, 200)
}

func TestCreateCustomReportWithKeepIPFilter(t *testing.T) {
	require := require.New(t)
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

	require.Equal(reportHeader, writer.CustomReportHeader)
	require.Equal(18, len(reportBody))

	for _, reportItem := range reportBody {
		require.Contains(reportItem[1], sampleIP)
	}
}

func TestCreateCustomReportWithKeepUAFilter(t *testing.T) {
	require := require.New(t)
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

	require.Equal(reportHeader, writer.CustomReportHeader)
	require.Equal(len(reportBody), 50)

	for _, reportItem := range reportBody {
		require.Contains(reportItem[5], sampleUserAgent)
	}
}

func TestCreateCustomReportWithUAFilter(t *testing.T) {
	require := require.New(t)
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

	require.Equal(reportHeader, writer.CustomReportHeader)

	for _, reportItem := range reportBody {
		require.NotContains(reportItem[5], sampleUserAgent)
	}
}

func TestCreateCustomReportWithKeepUriFilter(t *testing.T) {
	require := require.New(t)
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

	require.Equal(reportHeader, writer.CustomReportHeader)
	require.Equal(len(reportBody), 2)

	for _, reportItem := range reportBody {
		require.Contains(reportItem[2], sampleURI)
	}
}

func TestCreateCustomReportWithUriFilter(t *testing.T) {
	require := require.New(t)
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

	require.Equal(reportHeader, writer.CustomReportHeader)

	for _, reportItem := range reportBody {
		require.NotContains(reportItem[2], sampleURI)
	}
}

func TestCreateCustomReportWithMultipleAssetFilters(t *testing.T) {
	require := require.New(t)
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

	require.Equal(reportHeader, writer.CustomReportHeader)

	for _, reportItem := range reportBody {
		require.NotContains(reportItem[2], assetPattern1)
		require.NotContains(reportItem[2], assetPattern2)
	}
}

func TestCreateCustomReportWithKeepTimeFilter(t *testing.T) {
	require := require.New(t)
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

	require.Equal(reportHeader, writer.CustomReportHeader)

	timeFrameStart, _ := time.Parse(filter.EuropeFormat, testPeriodStart)
	timeFrameEnd, _ := time.Parse(filter.EuropeFormat, testPeriodFrameEnd)

	for _, reportItem := range reportBody {
		reportItemTime, _ := time.Parse(filter.EuropeFormat, reportItem[0])

		require.True(reportItemTime.After(timeFrameStart) && reportItemTime.Before(timeFrameEnd))
	}
}

func TestCreateCustomReportWithMultipleSystemIpProcessor(t *testing.T) {
	require := require.New(t)
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

	require.Equal(reportHeader, writer.CustomReportHeader)

	_, IPSubnet1, _ := net.ParseCIDR(systemIPSubnet1)

	for _, reportItem := range reportBody {
		ipList := strings.Split(reportItem[1], " ")

		for _, ip := range ipList {
			parsedIP := net.ParseIP(ip)

			require.NotEqual(systemIP2, ip)
			require.False(IPSubnet1.Contains(parsedIP))
		}
	}
}

func TestCreateCustomReportWithSubnetSystemIpProcessor(t *testing.T) {
	require := require.New(t)
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

	require.Equal(reportHeader, writer.CustomReportHeader)

	_, IPSubnet, _ := net.ParseCIDR(systemIPSubnet)

	for _, reportItem := range reportBody {
		ipList := strings.Split(reportItem[1], " ")

		for _, ip := range ipList {
			parsedIP := net.ParseIP(ip)

			require.False(IPSubnet.Contains(parsedIP))
		}
	}
}

func TestCreateCustomReportWithSingleSystemIpProcessor(t *testing.T) {
	require := require.New(t)
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

	require.Equal(reportHeader, writer.CustomReportHeader)

	for _, reportItem := range reportBody {
		ipList := strings.Split(reportItem[1], " ")

		for _, ip := range ipList {
			require.NotEqual(systemIP, ip)
		}
	}
}
