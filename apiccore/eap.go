package apiccore

import (
	"strings"
)

var (
	officialPrefixes map[string]struct{}
)

func init() {
	officialPrefixes = map[string]struct{}{
		"rlg-nrt-marketquest": {},
		"rlg-nrt-cbhome":      {},
		"rlg-nrt-trident":     {},
		"rlg-nrt-tridentods":  {},
		"rlg-rfg-dash":        {},
		"rlg-eduardo-data":    {},
	}
}

//IsOfficial check if the artifact belong to an official stack
func IsOfficial(artifactName string) bool {
	parts := strings.Split(artifactName, "-")
	if len(parts) >= 3 {
		artifactPrefix := strings.Join(parts[:3], "-")
		_, isOfficial := officialPrefixes[artifactPrefix]
		return isOfficial
	}
	return false
}
