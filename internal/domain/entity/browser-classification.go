package entity

//
type BrowserCategory struct {
	Agent    string
	Category string
}

//
func GetBrowserClassification() []BrowserCategory {
	return []BrowserCategory{
		{Agent: "Xbox One", Category: "Game Systems"},
		{Agent: "Xbox", Category: "Game Systems"},
		{Agent: "PlayStation", Category: "Game Systems"},
		{Agent: "NintendoBrowser", Category: "Game Systems"},
		{Agent: "Valve Steam", Category: "Game Systems"},
		{Agent: "Origin", Category: "Game Systems"},
		{Agent: "Raptr", Category: "Game Systems"},

		/* Based on Internet Explorer */
		{Agent: "America Online Browser", Category: "Others"},
		{Agent: "Avant Browser", Category: "Others"},

		/* Internet Explorer */
		{Agent: "IEMobile", Category: "MSIE"},
		{Agent: "MSIE", Category: "MSIE"},

		/* IE11 */
		{Agent: "Trident", Category: "MSIE"},

		/* Microsoft Edge */
		{Agent: "Edge", Category: "Edge"},

		/* Opera */
		{Agent: "Opera Mini", Category: "Opera"},
		{Agent: "Opera Mobi", Category: "Opera"},
		{Agent: "Opera", Category: "Opera"},
		{Agent: "OPR", Category: "Opera"},
		{Agent: "OPiOS", Category: "Opera"},
		{Agent: "Coast", Category: "Opera"},

		/* Payment Gateways */
		{Agent: "PayPal IPN", Category: "Payment Gateways"},

		/* Others */
		{Agent: "Homebrew", Category: "Others"},
		{Agent: "APT-HTTP", Category: "Others"},
		{Agent: "Apt-Cacher", Category: "Others"},
		{Agent: "Chef Client", Category: "Others"},
		{Agent: "Huawei", Category: "Others"},
		{Agent: "HUAWEI", Category: "Others"},
		{Agent: "BlackBerry", Category: "Others"},
		{Agent: "BrowserX", Category: "Others"},
		{Agent: "Dalvik", Category: "Others"},
		{Agent: "Dillo", Category: "Others"},
		{Agent: "ELinks", Category: "Others"},
		{Agent: "Epiphany", Category: "Others"},
		{Agent: "Firebird", Category: "Others"},
		{Agent: "Galeon", Category: "Others"},
		{Agent: "google-cloud-sdk", Category: "Others"},
		{Agent: "GranParadiso", Category: "Others"},
		{Agent: "IBrowse", Category: "Others"},
		{Agent: "K-Meleon", Category: "Others"},
		{Agent: "Kazehakase", Category: "Others"},
		{Agent: "Konqueror", Category: "Others"},
		{Agent: "Links", Category: "Others"},
		{Agent: "Lynx", Category: "Others"},
		{Agent: "Midori", Category: "Others"},
		{Agent: "Minefield", Category: "Others"},
		{Agent: "Mosaic", Category: "Others"},
		{Agent: "Netscape", Category: "Others"},
		{Agent: "SeaMonkey", Category: "Others"},
		{Agent: "UCBrowser", Category: "Others"},
		{Agent: "Wget", Category: "Others"},
		{Agent: "libfetch", Category: "Others"},
		{Agent: "check_http", Category: "Others"},
		{Agent: "Go-http-client", Category: "Others"},
		{Agent: "curl", Category: "Others"},
		{Agent: "midori", Category: "Others"},
		{Agent: "w3m", Category: "Others"},
		{Agent: "MicroMessenger", Category: "Others"},
		{Agent: "Apache", Category: "Others"},
		{Agent: "JOSM", Category: "Others"},

		/* Feed-reader-as-a-service */
		{Agent: "Feedspot", Category: "Feeds"},
		{Agent: "AppleNewsBot", Category: "Feeds"},
		{Agent: "Bloglines", Category: "Feeds"},
		{Agent: "Digg Feed Fetcher", Category: "Feeds"},
		{Agent: "Feedbin", Category: "Feeds"},
		{Agent: "FeedHQ", Category: "Feeds"},
		{Agent: "Feedly", Category: "Feeds"},
		{Agent: "Flipboard", Category: "Feeds"},
		{Agent: "Netvibes", Category: "Feeds"},
		{Agent: "NewsBlur", Category: "Feeds"},
		{Agent: "PinRSS", Category: "Feeds"},
		{Agent: "WordPress.com Reader", Category: "Feeds"},
		{Agent: "YandexBlogs", Category: "Feeds"},
		{Agent: "FeedFetcher-Google", Category: "Feeds"},
		{Agent: "startmebot", Category: "Feeds"},

		/* Google crawlers (some based on Chrome,
		 * therefore up on the list) */
		{Agent: "AdsBot-Google", Category: "Crawlers"},
		{Agent: "AppEngine-Google", Category: "Crawlers"},
		{Agent: "Mediapartners-Google", Category: "Crawlers"},
		{Agent: "Googlebot-Image", Category: "Crawlers"},
		{Agent: "Google Favicon", Category: "Crawlers"},
		{Agent: "Google", Category: "Crawlers"},
		{Agent: "WhatsApp", Category: "Crawlers"},

		/* Based on Firefox */
		{Agent: "Camino", Category: "Others"},

		/* Rebranded Firefox but is really unmodified
		* Firefox (Debian trademark policy) */
		{Agent: "Iceweasel", Category: "Firefox"},
		{Agent: "Focus", Category: "Firefox"},

		/* Klar is the name of Firefox Focus in the German market. */
		{Agent: "Klar", Category: "Firefox"},
		{Agent: "Firefox", Category: "Firefox"},

		/* Based on Chromium */
		{Agent: "BeakerBrowser", Category: "Beaker"},
		{Agent: "Brave", Category: "Brave"},
		{Agent: "Vivaldi", Category: "Vivaldi"},
		{Agent: "YaBrowser", Category: "Yandex.Browser"},

		/* Chrome has to go before Safari */
		{Agent: "HeadlessChrome", Category: "Chrome"},
		{Agent: "Chrome", Category: "Chrome"},
		{Agent: "CriOS", Category: "Chrome"},

		/* Crawlers/Bots (Possible Safari based) */
		{Agent: "AppleBot", Category: "Crawlers"},
		{Agent: "Twitter", Category: "Crawlers"},
		{Agent: "facebookexternalhit", Category: "Crawlers"},

		{Agent: "Safari", Category: "Safari"},

		/* Crawlers/Bots */
		{Agent: "BingPreview", Category: "Crawlers"},
		{Agent: "adidxbot", Category: "Crawlers"},
		{Agent: "bingbot", Category: "Crawlers"},
		{Agent: "Sogou", Category: "Crawlers"},
		{Agent: "Yahoo! Slurp", Category: "Crawlers"},
		{Agent: "Baiduspider", Category: "Crawlers"},
		{Agent: "Baidu", Category: "Crawlers"},
		{Agent: "Java", Category: "Crawlers"},
		{Agent: "Jakarta Commons-HttpClient", Category: "Crawlers"},
		{Agent: "netEstate", Category: "Crawlers"},
		{Agent: "PiplBot", Category: "Crawlers"},
		{Agent: "IstellaBot", Category: "Crawlers"},
		{Agent: "istellabot", Category: "Crawlers"},
		{Agent: "heritrix", Category: "Crawlers"},
		{Agent: "PagesInventory", Category: "Crawlers"},
		{Agent: "rogerbot", Category: "Crawlers"},
		{Agent: "fastbot", Category: "Crawlers"},
		{Agent: "yacybot", Category: "Crawlers"},
		{Agent: "PycURL", Category: "Crawlers"},
		{Agent: "PHP", Category: "Crawlers"},
		{Agent: "AndroidDownloadManager", Category: "Crawlers"},
		{Agent: "Embedly", Category: "Crawlers"},
		{Agent: "ruby", Category: "Crawlers"},
		{Agent: "Ruby", Category: "Crawlers"},
		{Agent: "python", Category: "Crawlers"},
		{Agent: "Python", Category: "Crawlers"},
		{Agent: "LinkedIn", Category: "Crawlers"},
		{Agent: "Microsoft-WebDAV", Category: "Crawlers"},
		{Agent: "GarlikCrawler", Category: "Crawlers"},
		{Agent: "CheckMarkNetwork", Category: "Crawlers"},
		{Agent: "uipbot", Category: "Crawlers"},
		{Agent: "FemtosearchBot", Category: "Crawlers"},
		{Agent: "ias_crawler", Category: "Crawlers"},
		{Agent: "Tbot", Category: "Crawlers"},
		{Agent: "CMS Crawler", Category: "Crawlers"},
		{Agent: "NetSeer crawler", Category: "Crawlers"},
		{Agent: "proximic", Category: "Crawlers"},
		{Agent: "ia_archiver", Category: "Crawlers"},
		{Agent: "DuckDuckBot", Category: "Crawlers"},
		{Agent: "Exabot", Category: "Crawlers"},
		{Agent: "HubSpot", Category: "Crawlers"},
		{Agent: "MJ12bot", Category: "Crawlers"},
		{Agent: "AhrefsBot", Category: "Crawlers"},
		{Agent: "DotBot", Category: "Crawlers"},
		{Agent: "SemrushBot", Category: "Crawlers"},
		{Agent: "YandexImages", Category: "Crawlers"},
		{Agent: "YandexBot", Category: "Crawlers"},
		{Agent: "msnbot-media", Category: "Crawlers"},
		{Agent: "msnbot", Category: "Crawlers"},
		{Agent: "VoilaBot", Category: "Crawlers"},
		{Agent: "Ask Jeeves/Teoma", Category: "Crawlers"},
		{Agent: "Pinterestbot", Category: "Crawlers"},
		{Agent: "Pinterest", Category: "Crawlers"},
		{Agent: "360Spider", Category: "Crawlers"},
		{Agent: "ZoominfoBot", Category: "Crawlers"},
		{Agent: "archive.org_bot", Category: "Crawlers"},
		{Agent: "Slackbot-LinkExpanding", Category: "Crawlers"},
		{Agent: "discobot", Category: "Crawlers"},
		{Agent: "BecomeBot", Category: "Crawlers"},
		{Agent: "ConveraCrawler", Category: "Crawlers"},
		{Agent: "GrapeshotCrawler", Category: "Crawlers"},
		{Agent: "Applebot", Category: "Crawlers"},

		/* Podcast fetchers */
		{Agent: "Downcast", Category: "Podcasts"},
		{Agent: "gPodder", Category: "Podcasts"},
		{Agent: "Instacast", Category: "Podcasts"},
		{Agent: "iTunes", Category: "Podcasts"},
		{Agent: "Miro", Category: "Podcasts"},
		{Agent: "Pocket Casts", Category: "Podcasts"},
		{Agent: "BashPodder", Category: "Podcasts"},

		/* Feed reader clients */
		{Agent: "Akregator", Category: "Feeds"},
		{Agent: "Apple-PubSub", Category: "Feeds"},
		{Agent: "com.apple.Safari.WebFeedParser", Category: "Feeds"},
		{Agent: "FeedDemon", Category: "Feeds"},
		{Agent: "Feedy", Category: "Feeds"},
		{Agent: "Liferea", Category: "Feeds"},
		{Agent: "NetNewsWire", Category: "Feeds"},
		{Agent: "RSSOwl", Category: "Feeds"},
		{Agent: "Thunderbird", Category: "Feeds"},

		{Agent: "Pingdom.com", Category: "Uptime"},
		{Agent: "jetmon", Category: "Uptime"},
		{Agent: "NodeUptime", Category: "Uptime"},
		{Agent: "NewRelicPinger", Category: "Uptime"},
		{Agent: "StatusCake", Category: "Uptime"},
		{Agent: "internetVista", Category: "Uptime"},
		{Agent: "Server Density Service Monitoring v2", Category: "Uptime"},
		{Agent: "Nexcess(tm) Site Performance Bot", Category: "Uptime"},
		{Agent: "UptimeRobot", Category: "Uptime"},

		/* Security Scanners */
		{Agent: "WPScan", Category: "Crawlers"},

		/* Services */
		{Agent: "okhttp", Category: "Crawlers"},

		{Agent: "Mozilla", Category: "Others"},
	}
}
