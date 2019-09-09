<p align="center">
  <img alt="Tango Logo" src="https://raw.githubusercontent.com/roma-glushko/tango/master/doc/tango-logo.png" height="120" />
  <h3 align="center">Tango</h3>
  <p align="center">"Dance" with your access logs</p>
</p>

---

<p align="center">
  
  [![Build Status](https://travis-ci.com/roma-glushko/tango.svg?branch=master)](https://travis-ci.com/roma-glushko/tango)
  [![Licence](https://img.shields.io/github/license/roma-glushko/tango)](https://github.com/roma-glushko/tango/blob/master/LICENSE)
  <img src="https://img.shields.io/badge/WIP-Work%20In%20Progress-yellow.svg" />
</p>

Tango is a dependency-free command-line tool for analyzing access logs 💃

Currently, work on this project is in progress. No stable or workable release yet.

## Installation

### macOS

Tango can be installed on macOS via <a href="https://brew.sh/">Homebrew</a>:

```
brew tap roma-glushko/tango
brew install roma-glushko/tango/tango
```

### Linux

Tango is available on <a href="https://snapcraft.io/tango">Snapcraft</a> via Snap.
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

### Windows

Tango can be installed on macOS via <a href="https://scoop.sh/">Scoop</a>:

```
scoop bucket add tango https://github.com/roma-glushko/scoop-tango.git
scoop install tango
```

## Use Cases

Legend:

- ✅ completed (MVP is ready)
- 👷 under development (will be available soon!)
- 🤔 no progress yet

### Reports

- Generate geo reports ✅
- Generate browser/crawler reports ✅
- Generate custom filtered reports ✅
- Generate visitor jorney pathes by IP ✅
- Generate request reports ✅
- Generate request frequency reports (automatically found automated requests from security scanners, exploits) 🤔
- Generate security reports (automatically found automated requests from security scanners, exploits) 🤔

### Filters

- Remove access log records from IP list ✅
- Keep only access log records from IP list ✅
- Filter access logs by user agent 🤔
- Extract access logs related to specific time frame ✅
- Ignore system IPs (like Harproxy, Fastly, any proxy, etc) ✅
- Filter access logs by response codes (keep only 30X, 50X, 40X responses, for example) 🤔
- Remove requests to web assets (js, css, images) ✅
- Remove other noisy requests that happens almost on the all pages (like customer data requests in Magento) ✅
- Find only records requested from the external resources 🤔
- Find only records requested during website browsing 🤔

### Output Formats

- Save reports to CSV files ✅
- Save reports to MySql database 🤔

### UX

- ability to set EU/US date time formats 🤔
- ability to generate output report filenames 🤔
- ability to show applied filters during report generation 🤔
- add hierarchical loading to separate base project configs from case-specific 🤔

## Usage

<p align="center">
    <img src="https://raw.githubusercontent.com/roma-glushko/tango/master/doc/tango.gif" />
</p>

TBD
