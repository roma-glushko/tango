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

### Global Filters

TBU

```bash
tango --uri-filter "/test-page" custom
tango --keep-uri-filter "/admin/" custom

tango --keep-time-filter "2019-09-15 04:16:00 -0400" --keep-time-filter "2019-09-15 04:35:00 -0400" custom

tango --ua-filter "iPhone OS 12_3_1 like Mac OS X" custom
tango --keep-ua-filter "iPhone OS 12_3_1 like Mac OS X" custom

tango --asset-filter "/pub/static/" --asset-filter "/pub/media/" custom
```

### Reports

#### Custom Reports

TBU

#### Geo Reports

TBU

#### Browser Reports

TBU

## Use Cases

List of usecases to cover: <a href="https://github.com/roma-glushko/tango/blob/master/doc/use-cases.md">Tango Usecases</a>
