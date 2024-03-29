package version

import (
	"github.com/mcuadros/go-version"
)

// Filter defines filter attributes for Versions.
type Filter struct {
	Version   string
	Minecraft string
}

// Versions is simply a collection of Version.
type Versions []Version

// Filter reduces the list of Versions by given criteria.
func (v Versions) Filter(filter *Filter) Versions {
	result := Versions{}

	vc := version.NewConstrainGroupFromString(filter.Version)
	mc := version.NewConstrainGroupFromString(filter.Minecraft)

	for _, row := range v {
		if filter.Version != "" {
			if !vc.Match(row.ID) {
				continue
			}
		}

		if filter.Minecraft != "" {
			if !mc.Match(row.Minecraft) {
				continue
			}
		}

		result = append(result, row)
	}

	return result
}
