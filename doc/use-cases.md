# Use Cases

Here is a list of usecases that we plan to cover in Tango.

Legend:

- ✅ completed (MVP is ready)
- 👷 under development (will be available soon!)
- 🤔 no progress yet

## Reports

- Generate geo reports ✅
- Generate browser/crawler reports ✅
- Generate custom filtered reports ✅
- Generate visitor jorney pathes by IP ✅
- Generate request reports ✅
- Generate request frequency reports (automatically found automated requests from security scanners, exploits) ✅
- Generate security reports (automatically found automated requests from security scanners, exploits) 🤔

## Filters

- Remove access log records from IP list ✅
- Keep only access log records from IP list ✅
- Filter access logs by user agent ✅
- Extract access logs related to specific time frame ✅
- Ignore system IPs (like Harproxy, Fastly, any proxy, etc) ✅
- Filter access logs by response codes (keep only 30X, 50X, 40X responses, for example) 🤔
- Remove requests to web assets (js, css, images) ✅
- Remove other noisy requests that happens almost on the all pages (like customer data requests in Magento) ✅
- Find only records requested from the external resources 🤔
- Find only records requested during website browsing 🤔

## Output Formats

- Save reports to CSV files ✅
- Save reports to MySql database 🤔

## UX

- ability to set EU/US date time formats 🤔
- ability to generate output report filenames 🤔
- ability to show applied filters during report generation 🤔
- add hierarchical loading to separate base project configs from case-specific 🤔

## Other

- Install/Reinstall automatically MaxMind Geo Library ✅