package version

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

const (
	// DefaultURL defines the default Forge version URL.
	DefaultURL = "https://files.minecraftforge.net/net/minecraftforge/forge/maven-metadata.json"
)

// FromString parses a version definition from string.
func FromString(content []byte) (Response, error) {
	result := Response{}

	if err := json.Unmarshal(content, &result); err != nil {
		return Response{}, errors.Wrap(err, "failed to parse versions")
	}

	return result, nil
}

// FromURL parses a version definition from URL.
func FromURL(path string) (Response, error) {
	resp, err := http.Get(path)

	if err != nil {
		return Response{}, errors.Wrap(err, "failed to fetch versions")
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return Response{}, errors.Wrap(err, "failed to read versions")
	}

	return FromString(body)
}

// FromPath parses a version definition from file path.
func FromPath(path string) (Response, error) {
	body, err := os.ReadFile(path)

	if err != nil {
		return Response{}, errors.Wrap(err, "failed to read versions")
	}

	return FromString(body)
}

// FromDefault is a simply wrapper that loads the default URL.
func FromDefault() (Response, error) {
	return FromURL(DefaultURL)
}
