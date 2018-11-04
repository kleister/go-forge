package version

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kleister/go-forge/version/internal"
	"github.com/mcuadros/go-version"
)

var (
	// OldestMinecraft defines the oldest allowed Minecraft version.
	OldestMinecraft = "1.6.4"
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

	for _, row := range result.Number {
		if version.Compare(row.Mcversion, OldestMinecraft, ">") {
			v := Version{
				ID:        row.Version,
				Minecraft: row.Mcversion,
				URL: fmt.Sprintf(
					"%s%s-%s/forge-%s-%s-universal.jar",
					result.Homepage,
					row.Mcversion,
					row.Version,
					row.Mcversion,
					row.Version,
				),
				Release: time.Unix(int64(row.Modified), 0),
			}

			if val, ok := result.Promos[fmt.Sprintf("%s-recommended", row.Mcversion)]; ok {
				v.Recommended = row.Build == val
			}

			if val, ok := result.Promos[fmt.Sprintf("%s-latest", row.Mcversion)]; ok {
				v.Latest = row.Build == val
			}

			r.Releases = append(r.Releases, v)
		}
	}

	return nil
}
