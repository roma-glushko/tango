package test

import (
	"os"
	"tango/pkg/adapters/writer"
	"tango/pkg/cli"
	"tango/pkg/di"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInstallGeoLib(t *testing.T) {
	require := require.New(t)
	tangoCli := cli.NewTangoCli("0.0.0-test", "dummycommithash")

	tangoCli.Run([]string{
		"main",
		"geo-lib",
		"-a",
		"197343",
		"-l",
		"aD36bIADbkErTRrS",
	})

	_, geoConfExistErr := di.InitMaxmindConfResolver().GetPath()
	require.False(os.IsNotExist(geoConfExistErr))

	_, geoLibExistErr := di.InitMaxmindGeoLibResolver().GetPath()
	require.False(os.IsNotExist(geoLibExistErr))
}

func TestCreateGeoReportWithSystemIpProcessor(t *testing.T) {
	require := require.New(t)
	tangoCli := cli.NewTangoCli("0.0.0-test", "dummycommithash")

	reportFilePath := "results/geo-report-with-system-ip-processor.csv"

	tangoCli.Run([]string{
		"main",
		"-c",
		"fixture/.tango.system-ips.yaml",
		"geo",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	reportHeader, reportBody := testReport[0], testReport[1:]

	require.Equal(reportHeader, writer.GeoReportHeader)
	require.Len(reportBody, 40)

	testGeoData := map[string]struct {
		Country   string
		City      string
		Continent string
		Requests  string
	}{
		"130.93.253.236": {
			Country:   "Hungary",
			City:      "",
			Continent: "Europe",
			Requests:  "18",
		},
		"121.79.80.29": {
			Country:   "Australia",
			City:      "Brisbane",
			Continent: "Oceania",
			Requests:  "4",
		},
	}

	testSystemIPList := []string{
		"157.52.99.32",
		"157.52.99.35",
		"157.52.99.37",
		"157.52.99.44",
		"157.52.75.66",
		"157.52.67.41",
		"157.52.75.43",
		"104.156.87.35",
		"199.27.79.24",
		"199.27.79.25",
		"104.156.91.44",
	}

	for _, reportItem := range reportBody {
		ip := reportItem[0]

		if expected, ok := testGeoData[ip]; ok {
			require.Equal(expected.Country, reportItem[1])
			require.Equal(expected.City, reportItem[2])
			require.Equal(expected.Continent, reportItem[3])
			require.NotEmpty(reportItem[4]) // Sample Request (order-dependent)
			require.NotEmpty(reportItem[5]) // Browser Agent (order-dependent)
			require.Equal(expected.Requests, reportItem[6])
		}

		require.NotContains(testSystemIPList, ip)
	}
}
