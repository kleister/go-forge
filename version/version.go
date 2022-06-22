package version

// Version defines the standard version format.
type Version struct {
	ID        string `json:"id"`
	Minecraft string `json:"minecraft"`
	URL       string `json:"url"`
}
