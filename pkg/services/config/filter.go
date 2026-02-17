package config

// FilterConfig holds configuration for access log filtering.
type FilterConfig struct {
	AssetFilters   []string
	KeepTimeFrames []string
	URIFilters     []string
	KeepURIFilters []string
	IPFilters      []string
	KeepIPFilters  []string
	UaFilters      []string
	KeepUaFilters  []string
}

// NewFilterConfig creates a new FilterConfig with the given parameters.
func NewFilterConfig(
	assetFilters []string,
	keepTimeFrames []string,
	uriFilters []string,
	keepURIFilters []string,
	ipFilters []string,
	keepIPFilters []string,
	uaFilters []string,
	keepUaFilters []string,
) FilterConfig {
	return FilterConfig{
		AssetFilters:   assetFilters,
		KeepTimeFrames: keepTimeFrames,
		URIFilters:     uriFilters,
		KeepURIFilters: keepURIFilters,
		IPFilters:      ipFilters,
		KeepIPFilters:  keepIPFilters,
		UaFilters:      uaFilters,
		KeepUaFilters:  keepUaFilters,
	}
}
