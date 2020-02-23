<p align="center">
  <img alt="Tango" src="https://raw.githubusercontent.com/roma-glushko/tango/master/doc/tango-logo.png" height="90" />
  <h3 align="center">Tango</h3>
  <p align="center">Tool to get insights from the server access logs</p>
</p>

---

<p align="center">
  <a href="https://travis-ci.org/roma-glushko/tango" alt="Build Status"><img alt="Tango" src="https://travis-ci.org/roma-glushko/tango.svg?branch=master" /></a>
  <a href="https://github.com/roma-glushko/tango/blob/master/LICENSE" alt="License"><img alt="License" src="https://img.shields.io/github/license/roma-glushko/tango" /></a>
  <img src="https://img.shields.io/badge/WIP-Work%20In%20Progress-yellow.svg" />
</p>

<p align="center">
    <img src="https://raw.githubusercontent.com/roma-glushko/tango/master/doc/tango.gif" width="500px" />
</p>

Tango is a dependency-free command-line tool for analyzing access logs 💃

Currently, work on this project is in progress. 
However, a few pre-releases are ready available to use 🎉

## Installation

### macOS

<p align="center">	
  <img src="https://raw.githubusercontent.com/roma-glushko/tango/master/doc/tango-install-homebrew.gif" width="500px" />
</p>

Tango can be installed on macOS via <a href="https://brew.sh/">Homebrew</a>:

```bash
brew tap roma-glushko/tango
brew install roma-glushko/tango/tango
```

To upgrade, try to run:

```bash
brew upgrade tango
```

### Linux

Tango is available on Linux via <a href="https://snapcraft.io/tango">Snapcraft</a>.
This means that Tango can be installed on:

- <a href="https://snapcraft.io/install/tango/ubuntu">Ubuntu</a>
- <a href="https://snapcraft.io/install/tango/debian">Debian</a>
- <a href="https://snapcraft.io/install/tango/centos">CentOS</a>
- <a href="https://snapcraft.io/install/tango/opensuse">openSUSE</a>
- <a href="https://snapcraft.io/install/tango/mint">Linux Mint</a>
- <a href="https://snapcraft.io/install/tango/fedora">Fedora</a>
- <a href="https://snapcraft.io/install/tango/kubuntu">Kubuntu</a>
- <a href="https://snapcraft.io/install/tango/elementary">elementary OS</a>
- <a href="https://snapcraft.io/install/tango/arch">Arch Linux</a>
- <a href="https://snapcraft.io/install/tango/kde-neon">KDE Neon</a>
- <a href="https://snapcraft.io/install/tango/manjaro">Manjaro</a>

To upgrade, try to run:

```bash
snap refresh tango
```

### Windows

Tango can be installed on Windows via <a href="https://scoop.sh/">Scoop</a>:

```bash
scoop bucket add tango https://github.com/roma-glushko/scoop-tango.git
scoop install tango
```

To upgrade, try to run:

```bash
scoop update tango
```

## Usage

List of available commands:

```bash
tango help
```

Tango Version:

```bash
tango -v
```

### Global Options

#### Filters

```bash
// IP filters
tango --ip-filter "127.0.0.1" custom -l access-log.log -r custom.csv
tango --keep-ip-filter "8.8.8.8" custom -l access-log.log -r custom.csv
```

```bash
// URI filters
tango --uri-filter "/test-page" custom -l access-log.log -r custom.csv
tango --keep-uri-filter "/admin/" custom -l access-log.log -r custom.csv
```

```bash
// Time Frame filter
tango --keep-time-filter "2019-09-15 04:16:00 -0400" --keep-time-filter "2019-09-15 04:35:00 -0400" custom -l access-log.log -r custom.csv
```

```bash
// User Agent filters
tango --ua-filter "iPhone OS 12_3_1 like Mac OS X" custom -l access-log.log -r custom.csv
tango --keep-ua-filter "iPhone OS 12_3_1 like Mac OS X" custom -l access-log.log -r custom.csv
```

```bash
// Asset filter
tango --asset-filter "/pub/static/" --asset-filter "/pub/media/" custom -l access-log.log -r custom.csv
```

```bash
// System IP filter
tango --system-ips "127.0.0.1"  --system-ips "1.2.3.4" custom -l access-log.log -r custom.csv
```

#### Other

```bash
// Base URL info
tango --base-url "https://example.com/" custom -l access-log.log -r custom.csv
```

### Report Commands

#### Custom Reports

```bash
tango --keep-uri-filter "/newsletter/subscriber/new/" custom -l access-log.log -r custom.csv
```

Use cases:

- generate a report with all requests from a certain IP
- generate a report with all requests to a certain URL

#### Geo Reports

```bash
tango geo -l access-log.log -r custom.csv
```

Use cases:

- collects geo information about all IPs that requested the website
- get request distribution by IP with geo information
- see all IPs sorted by countries/continents/cities  

Example of the report:

<details>
  <summary>Example of the report</summary>
  
  | IP             | Country       | City    | Continent     | Sample Request | Browser Agent                                                            | Count of Requests |
|----------------|---------------|---------|---------------|----------------|--------------------------------------------------------------------------|-------------------|
| 46.229.173.68  | United States | Ashburn | North America | /robots.txt    | Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html) | 362               |
| 40.77.167.91   | United States | Boydton | North America | /contact-us    | Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)  | 3                 |
| 178.154.171.62 | Russia        |         | Europe        | /              | Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)         | 34                |
  
</details>

#### Browser Reports

```bash
tango browser -l access-log.log -r custom.csv
```

Use cases:

- check how many requests were sent by crawlers
- check what kind of browsers requested the website
- check bandwith that was transmitted to all kind of browsers
- check what crawlers requested the website

<details>
  <summary>Example of the report</summary>
  
  | Category | Browser | Requests | Bandwith | Sample URL | User Agents |
|----------|---------|----------|----------|--------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Crawlers | bingbot | 629 | 28.8 MB | /black-bag-product | Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm) |
| Chrome | Chrome | 131998 | 1.3 GB | /gears/bags?p=3 | Mozilla/5.0 (Linux; Android 8.0.0; G8441) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36<br>Mozilla/5.0 (Linux; Android 9; SM-G960F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 MobileSafari/537.36<br>Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36<br>Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.85 Safari/537.36 |
  
</details>

#### Request Reports

```bash
tango request -l access-log.log -r custom.csv
```

Use cases:

- check how many requests were sent to a certain URL
- check all URLs that were responded with 404/50X code
- find requests from security scanners (sort by response codes and look at 404/50X codes which were requested only 1 time)

<details>
  <summary>Example of the report</summary>
  
  | Path | Requests | Response Code | Referer URLs |
|---------------------------------------|----------|---------------|---------------------------------------|
| /media/catalog/product/black-bag.jpg | 20 | 200 | /black-bag |
| /admin/sales/order/view/order_id/1234 | 4 | 200 | /admin/sales/order/index/order_id/123 |
| /test321 | 1 | 404 | / |
  
</details>

#### Pace Reports [Experimental]

```bash
tango pace -l access-log.log -r custom.csv
```

Use cases:

- check which IPs and how many requests they made during a certain time frame
- check count of requests per minutes/hours

<details>
  <summary>Example of the report</summary>
  
  | Hour Group | Minute Group | IP | Browser | Pace (req/min) | Pace (req/hour) |
|-----------------|------------------|---------------|--------------------------------------------------------------------|----------------|-----------------|
| 2020-02-10 04 h |  |  |  |  | 35 |
|  | 2020-02-10 04:06 |  |  | 15 |  |
|  |  | 51.15.191.180 | Barkrowler/0.9 (+https://babbar.tech/crawler) | 10 |  |
|  |  | 54.36.150.167 | Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/) | 5 |  |
|  | 2020-02-10 04:06 |  |  | 15 |  |
|  | 2020-02-10 04:07 |  |  | 20 |  |
|  |  | 66.249.76.89 | Googlebot-Image/1.0 | 20 |  |
|  | 2020-02-10 04:07 |  |  | 20 |  |
| 2020-02-10 04 h |  |  |  |  | 35 |
  
</details>

#### Journey Reports [Experimental]

```bash
tango journey -l access-log.log -r custom.csv
```

### Misc Commands

```bash
// Install geo library to get more info in geo reports
tango geo-lib
```

Tango uses the MaxMind GeoLite2-City database and stores it under: 

- macOS - `/Users/[username]/.tango/GeoLite2-City.mmdb`

### Example of the config file

Put the similar content to a `.tango.yaml` file under your working directory where you analyze logs: 

```yaml
"asset-filter":
  - "/pub/static/"
  - "/pub/media/"
"ip-filter":
  - "127.0.0.1"
"system-ips":
  # Fastly IPs
  - "23.235.32.0/20"
  - "43.249.72.0/22"
  - "103.244.50.0/24"
  - "103.245.222.0/23"
  - "103.245.224.0/24"
  - "104.156.80.0/20"
  - "151.101.0.0/16"
  - "157.52.64.0/18"
  - "167.82.0.0/17"
  - "167.82.128.0/20"
  - "167.82.160.0/20"
  - "167.82.224.0/20"
  - "172.111.64.0/18"
  - "185.31.16.0/22"
  - "199.27.72.0/21"
  - "199.232.0.0/16"
```

## Use Cases

List of usecases to cover: <a href="https://github.com/roma-glushko/tango/blob/master/doc/use-cases.md">Tango Usecases</a>
