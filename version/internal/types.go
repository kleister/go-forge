package internal

// Root defines the root response element for forge versions.
type Root map[string][]string

// Slim defines the root response element for the promotions.
type Slim struct {
	Homepage string            `json:"homepage"`
	Promos   map[string]string `json:"promos"`
}
