package version

import (
	"fmt"
	"time"

	"github.com/json-iterator/go"
	"github.com/kleister/go-forge/version/internal"
)

// Response is the standard root element for the version list.
type Response struct {
	Releases Versions `json:"releases"`
}

// UnmarshalJSON implements the custom JSON unmarshaler.
func (r *Response) UnmarshalJSON(b []byte) error {
	result := internal.Root{}

	if err := jsoniter.Unmarshal(b, &result); err != nil {
		return err
	}

	for _, number := range result.Number {
		v := Version{
			ID:        number.Version,
			Minecraft: number.Mcversion,
			URL: fmt.Sprintf(
				"%s%s-%s/forge-%s-%s-universal.jar",
				result.Homepage,
				number.Mcversion,
				number.Version,
				number.Mcversion,
				number.Version,
			),
			Release: time.Unix(int64(number.Modified), 0),
		}

		if val, ok := result.Promos[fmt.Sprintf("%s-recommended", number.Mcversion)]; ok {
			v.Recommended = number.Build == val
		}

		if val, ok := result.Promos[fmt.Sprintf("%s-latest", number.Mcversion)]; ok {
			v.Latest = number.Build == val
		}

		r.Releases = append(r.Releases, v)
	}

	return nil
}
