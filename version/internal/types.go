package internal

// Version defines the response version.
type Version struct {
	Build     int
	Mcversion string
	Modified  float64
	Version   string
}

// Root defines the root response element.
type Root struct {
	Homepage string
	Number   map[string]Version
	Promos   map[string]int
}
