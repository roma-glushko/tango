package entity

//
func GetBrowserClassification() map[string]string {
	return map[string]string{
		"Xbox One":        "Game Systems",
		"Xbox":            "Game Systems",
		"PlayStation":     "Game Systems",
		"NintendoBrowser": "Game Systems",
		"Valve Steam":     "Game Systems",
		"Origin":          "Game Systems",
		"Raptr":           "Game Systems",

		/* Based on Internet Explorer */
		"America Online Browser": "Others",
		"Avant Browser":          "Others",

		/* Internet Explorer */
		"IEMobile": "MSIE",
		"MSIE":     "MSIE",

		/* IE11 */
		"Trident/7.0": "MSIE",

		/* Microsoft Edge */
		"Edge": "Edge",

		/* Opera */
		"Opera Mini": "Opera",
		"Opera Mobi": "Opera",
		"Opera":      "Opera",
		"OPR":        "Opera",
		"OPiOS":      "Opera",
		"Coast":      "Opera",

		/* Payment Gateways */
		"PayPal IPN": "Payment Gateways ",

		/* Others */
		"Homebrew":         "Others",
		"APT-HTTP":         "Others",
		"Apt-Cacher":       "Others",
		"Chef Client":      "Others",
		"Huawei":           "Others",
		"HUAWEI":           "Others",
		"BlackBerry":       "Others",
		"BrowserX":         "Others",
		"Dalvik":           "Others",
		"Dillo":            "Others",
		"ELinks":           "Others",
		"Epiphany":         "Others",
		"Firebird":         "Others",
		"Galeon":           "Others",
		"google-cloud-sdk": "Others",
		"GranParadiso":     "Others",
		"IBrowse":          "Others",
		"K-Meleon":         "Others",
		"Kazehakase":       "Others",
		"Konqueror":        "Others",
		"Links":            "Others",
		"Lynx":             "Others",
		"Midori":           "Others",
		"Minefield":        "Others",
		"Mosaic":           "Others",
		"Netscape":         "Others",
		"SeaMonkey":        "Others",
		"UCBrowser":        "Others",
		"Wget":             "Others",
		"libfetch":         "Others",
		"check_http":       "Others",
		"Go-http-client":   "Others",
		"curl":             "Others",
		"midori":           "Others",
		"w3m":              "Others",
		"MicroMessenger":   "Others",
		"Apache":           "Others",
		"JOSM":             "Others",

		/* Feed-reader-as-a-service */
		"AppleNewsBot":         "Feeds",
		"Bloglines":            "Feeds",
		"Digg Feed Fetcher":    "Feeds",
		"Feedbin":              "Feeds",
		"FeedHQ":               "Feeds",
		"Feedly":               "Feeds",
		"Flipboard":            "Feeds",
		"Netvibes":             "Feeds",
		"NewsBlur":             "Feeds",
		"PinRSS":               "Feeds",
		"WordPress.com Reader": "Feeds",
		"YandexBlogs":          "Feeds",

		/* Google crawlers (some based on Chrome,
		 * therefore up on the list) */
		"AdsBot-Google":        "Crawlers",
		"AppEngine-Google":     "Crawlers",
		"Mediapartners-Google": "Crawlers",
		"Google":               "Crawlers",
		"WhatsApp":             "Crawlers",

		/* Based on Firefox */
		"Camino": "Others",
		/* Rebranded Firefox but is really unmodified
		* Firefox (Debian trademark policy) */
		"Iceweasel": "Firefox",
		"Focus":     "Firefox",
		/* Klar is the name of Firefox Focus in the German market. */
		"Klar":    "Firefox",
		"Firefox": "Firefox",

		/* Based on Chromium */
		"BeakerBrowser": "Beaker",
		"Brave":         "Brave",
		"Vivaldi":       "Vivaldi",
		"YaBrowser":     "Yandex.Browser",

		/* Chrome has to go before Safari */
		"HeadlessChrome": "Chrome",
		"Chrome":         "Chrome",
		"CriOS":          "Chrome",

		/* Crawlers/Bots (Possible Safari based) */
		"AppleBot":            "Crawlers",
		"Twitter":             "Crawlers",
		"facebookexternalhit": "Crawlers",

		"Safari": "Safari",

		/* Crawlers/Bots */
		"Sogou":                      "Crawlers",
		"Baidu":                      "Crawlers",
		"Java":                       "Crawlers",
		"Jakarta Commons-HttpClient": "Crawlers",
		"netEstate":                  "Crawlers",
		"PiplBot":                    "Crawlers",
		"IstellaBot":                 "Crawlers",
		"heritrix":                   "Crawlers",
		"PagesInventory":             "Crawlers",
		"rogerbot":                   "Crawlers",
		"fastbot":                    "Crawlers",
		"yacybot":                    "Crawlers",
		"PycURL":                     "Crawlers",
		"PHP":                        "Crawlers",
		"AndroidDownloadManager":     "Crawlers",
		"Embedly":                    "Crawlers",
		"ruby":                       "Crawlers",
		"Ruby":                       "Crawlers",
		"python":                     "Crawlers",
		"Python":                     "Crawlers",
		"LinkedIn":                   "Crawlers",
		"Microsoft-WebDAV":           "Crawlers",
		"bingbot":                    "Crawlers",

		/* Podcast fetchers */
		"Downcast":     "Podcasts",
		"gPodder":      "Podcasts",
		"Instacast":    "Podcasts",
		"iTunes":       "Podcasts",
		"Miro":         "Podcasts",
		"Pocket Casts": "Podcasts",
		"BashPodder":   "Podcasts",

		/* Feed reader clients */
		"Akregator":                      "Feeds",
		"Apple-PubSub":                   "Feeds",
		"com.apple.Safari.WebFeedParser": "Feeds",
		"FeedDemon":                      "Feeds",
		"Feedy":                          "Feeds",
		"Liferea":                        "Feeds",
		"NetNewsWire":                    "Feeds",
		"RSSOwl":                         "Feeds",
		"Thunderbird":                    "Feeds",

		"Pingdom.com":                          "Uptime",
		"jetmon":                               "Uptime",
		"NodeUptime":                           "Uptime",
		"NewRelicPinger":                       "Uptime",
		"StatusCake":                           "Uptime",
		"internetVista":                        "Uptime",
		"Server Density Service Monitoring v2": "Uptime",
		"Nexcess(tm) Site Performance Bot":     "Uptime",

		/* Security Scanners */
		"WPScan": "Crawlers",

		/* Services */
		"okhttp": "Crawlers",

		"Mozilla": "Others",
	}
}
