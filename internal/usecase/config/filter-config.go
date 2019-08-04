package config

//
type FilterConfig struct {
	AssetFilters   []string
	KeepTimeFrames []string
	UriFilters     []string
	KeepUriFilters []string
	IpFilters      []string
	KeepIpFilters  []string
	UaFilter       string
	KeepUaFilter   string
}

//
func NewFilterConfig(
	assetFilters []string,
	keepTimeFrames []string,
	uriFilters []string,
	keepUriFilters []string,
	ipFilters []string,
	keepIpFilters []string,
	uaFilter string,
	keepUaFilter string,
) FilterConfig {
	return FilterConfig{
		AssetFilters:   assetFilters,
		KeepTimeFrames: keepTimeFrames,
		UriFilters:     uriFilters,
		KeepUriFilters: keepUriFilters,
		IpFilters:      ipFilters,
		KeepIpFilters:  keepIpFilters,
		UaFilter:       uaFilter,
		KeepUaFilter:   keepUaFilter,
	}
}
