package test

import (
	"tango/pkg/adapters/writer"
	"tango/pkg/cli"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateBrowserReport(t *testing.T) {
	require := require.New(t)
	tangoCli := cli.NewTangoCli("0.0.0-test", "dummycommithash")

	reportFilePath := "results/browser-report.csv"

	tangoCli.Run([]string{
		"main",
		"-c",
		"fixture/.tango.empty.yaml",
		"browser",
		"-l",
		"fixture/apache-combined-access-log-jul-200rec-with-timezone.log",
		"-r",
		reportFilePath,
	})

	testReport := GetTestCsvReport(reportFilePath, t)

	reportHeader, reportBody := testReport[0], testReport[1:]

	require.Equal(reportHeader, writer.BrowserReportHeader)

	testBrowserData := map[string]map[string][]string{
		"Crawlers-bingbot": {
			"general": {
				"Crawlers",          // Category
				"bingbot",           // Browser
				"2",                 // Request
				"200.5 kB",          // Bandwidth
				"/category200?p=22", // Requests
			},
			"browsers": {
				"Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)", // Full Browser Agent
			},
		},
		"Crawlers-Googlebot": {
			"general": {
				"Crawlers",                   // Category
				"Googlebot",                  // Browser
				"10",                         // Requests
				"797.5 kB",                   // Bandwidth
				"/product-dsm88-shmoav.html", // Request
			},
			"browsers": {
				"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
				"Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			},
		},
	}

	for _, reportItem := range reportBody {
		browserCategory := reportItem[0]
		browserName := reportItem[1]
		browserAgents := reportItem[5]

		testDataKey := browserCategory + "-" + browserName

		if expectedItem, ok := testBrowserData[testDataKey]; ok {
			generalData := reportItem[:len(reportItem)-1]

			require.ElementsMatch(expectedItem["general"], generalData)

			for _, browserSample := range expectedItem["browsers"] {
				require.Contains(browserAgents, browserSample)
			}
		}
	}
}
