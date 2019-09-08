<p align="center">
  <img alt="Tango Logo" src="https://raw.githubusercontent.com/roma-glushko/tango/master/doc/tango-logo.png" height="120" />
  <h3 align="center">Tango</h3>
  <p align="center">"Dance" with your access logs</p>
</p>

---

<img src="https://raw.githubusercontent.com/roma-glushko/tango/master/doc/tango.gif" />

Tango is a command-line tool for "dancing" with access logs 💃

Currently, work on this project is in progress. No stable or workable release yet.

## Badges

<img src="https://travis-ci.com/roma-glushko/tango.svg?branch=master" />
<img src="https://img.shields.io/badge/WIP-Work%20In%20Progress-yellow.svg" />

## Use Cases

Legend:

- ✅ completed (MVP is ready)
- 👷 under development (will be available soon!)
- 🤔 no progress so far

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

TBD
