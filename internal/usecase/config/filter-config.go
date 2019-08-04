package config

//
type FilterConfig struct {
	AssetFilters   []string
	KeepTimeFrames []string
	UriFilters     []string
	KeepUriFilters []string
	IpFilters      []string
	KeepIpFilters  []string
	UaFilters      []string
	KeepUaFilters  []string
}

//
func NewFilterConfig(
	assetFilters []string,
	keepTimeFrames []string,
	uriFilters []string,
	keepUriFilters []string,
	ipFilters []string,
	keepIpFilters []string,
	uaFilters []string,
	keepUaFilters []string,
) FilterConfig {
	return FilterConfig{
		AssetFilters:   assetFilters,
		KeepTimeFrames: keepTimeFrames,
		UriFilters:     uriFilters,
		KeepUriFilters: keepUriFilters,
		IpFilters:      ipFilters,
		KeepIpFilters:  keepIpFilters,
		UaFilters:      uaFilters,
		KeepUaFilters:  keepUaFilters,
	}
}
