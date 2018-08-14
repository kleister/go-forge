package version

import (
	"time"
)

// Version defines the standard version format.
type Version struct {
	ID          string    `json:"id"`
	Minecraft   string    `json:"minecraft"`
	URL         string    `json:"url"`
	Recommended bool      `json:"recommended"`
	Latest      bool      `json:"latest"`
	Release     time.Time `json:"release"`
}
