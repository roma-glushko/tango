package filter

import (
	"strings"
)

var webAssetPatterns = []string{
	"/pub/static/",
	"/pub/media/",
}

//
func FilterWebAsset(uri string) bool {
	for _, webAssetPattern := range webAssetPatterns {
		if strings.HasPrefix(uri, webAssetPattern) {
			return true
		}
	}

	return false
}
