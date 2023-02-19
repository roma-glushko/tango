package test

import (
	"os"
	"tango/pkg/adapters/writer"
	"tango/pkg/cli"
	"tango/pkg/di"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstallGeoLib(t *testing.T) {
	asserter := assert.New(t)
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
	asserter.False(os.IsNotExist(geoConfExistErr), "MaxMind Configuration File was not created")

	_, geoLibExistErr := di.InitMaxmindGeoLibResolver().GetPath()
	asserter.False(os.IsNotExist(geoLibExistErr), "MaxMind Geo Lib was not created")
}

func TestCreateGeoReportWithSystemIpProcessor(t *testing.T) {
	asserter := assert.New(t)
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

	asserter.Equal(reportHeader, writer.GeoReportHeader, "Geo Report Header is different")
	asserter.Len(reportBody, 40, "Geo Report should contain 40 record")

	testGeoData := map[string][]string{
		"130.93.253.236": {
			"130.93.253.236",  // IP
			"Hungary",         // Country
			"Szentlorinckata", // City
			"",                // Sample Request, keep it empty as it may change since it's a random sampled URL
			"Europe",          // Continent
			"Mozilla/5.0 (iPhone; CPU iPhone OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.1 Mobile/15E148 Safari/604.1", //Browser Agent
			"18", // Count of Requests
		},
		"121.79.80.29": {
			"121.79.80.29", // IP
			"Australia",    // Country
			"",             // City
			"",             // Sample Request, keep it empty as it may change since it's a random sampled URL
			"Oceania",      // Continent
			"Mozilla/5.0 (X11; CrOS x86_64 12105.75.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.102 Safari/537.36", //Browser Agent
			"4", // Count of Requests
		},
	}

	testSystemIpList := []string{
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
		IP := reportItem[0]

		if expectedItem, found := testGeoData[IP]; found {
			reportItem[4] = "" // ignore the randomly sampled URL

			asserter.ElementsMatch(expectedItem, reportItem, "Sample Geo Report Item should match")
		}

		asserter.NotContains(testSystemIpList, IP, "System IP should be filtered")
	}

}
