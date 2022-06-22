package version

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kleister/go-forge/version/internal"
	"github.com/mcuadros/go-version"
)

var (
	// OldestMinecraft defines the oldest allowed Minecraft version.
	OldestMinecraft = "1.6.4"

	// DownloadHomepage defines the download base url to fetch installer.
	DownloadHomepage = "https://maven.minecraftforge.net/net/minecraftforge/forge"
)

// Response is the standard root element for the version list.
type Response struct {
	Releases Versions `json:"releases"`
}

// UnmarshalJSON implements the custom JSON unmarshaler.
func (r *Response) UnmarshalJSON(b []byte) error {
	result := internal.Root{}

	if err := json.Unmarshal(b, &result); err != nil {
		return err
	}

	for minecraft, versions := range result {
		if version.Compare(minecraft, OldestMinecraft, "<=") {
			continue
		}

		for _, version := range versions {
			split := strings.Split(version, "-")

			v := Version{
				ID:        split[1],
				Minecraft: minecraft,
				URL: fmt.Sprintf(
					"%s/%s/forge-%s-universal.jar",
					DownloadHomepage,
					version,
					version,
				),
			}

			r.Releases = append(r.Releases, v)
		}
	}

	return nil
}
