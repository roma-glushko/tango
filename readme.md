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

#### Geo Reports

```bash
tango geo -l access-log.log -r custom.csv
```

#### Browser Reports

```bash
tango browser -l access-log.log -r custom.csv
```

#### Request Reports

```bash
tango request -l access-log.log -r custom.csv
```

#### Pace Reports [Experimental]

```bash
tango pace -l access-log.log -r custom.csv
```

#### Journey Reports [Experimental]

```bash
tango journey -l access-log.log -r custom.csv
```

### Misc Commands

```bash
// Install geo library to get more info in geo reports
tango geo-lib
```

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
