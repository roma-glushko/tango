package test

import (
	"tango/pkg/cli"
	"tango/pkg/infrastructure/writer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBrowserReport(t *testing.T) {
	assert := assert.New(t)
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

	assert.Equal(reportHeader, writer.BrowserReportHeader, "Browser Report Header is different")
	//assert.Len(reportBody, 40, "Geo Report should contain 40 record")

	testBrowserData := map[string]map[string][]string{
		"Crawlers-bingbot": map[string][]string{
			"general": []string{
				"Crawlers",          // Category
				"bingbot",           // Browser
				"2",                 // Request
				"200.5 kB",          // Bandwith
				"/category200?p=22", // Requests
			},
			"browsers": []string{
				"Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)", // Full Browser Agent
			},
		},
		"Crawlers-Googlebot": map[string][]string{
			"general": []string{
				"Crawlers",                   // Category
				"Googlebot",                  // Browser
				"10",                         // Requests
				"797.5 kB",                   // Bandwith
				"/product-dsm88-shmoav.html", // Request
			},
			"browsers": []string{
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

			assert.ElementsMatch(expectedItem["general"], generalData, "Browser Report Items should match samples")

			for _, browserSample := range expectedItem["browsers"] {
				assert.Contains(browserAgents, browserSample, "Browser Report Items should match samples")
			}
		}
	}
}
