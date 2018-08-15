package version

import (
	"github.com/mcuadros/go-version"
	"sort"
)

// ByVersion sorts a list of Forge versions by ID.
type ByVersion Versions

// Sort simply sorts the versions list.
func (b ByVersion) Sort() {
	sort.Sort(b)
}

// Len is part of the sorting algorithm.
func (b ByVersion) Len() int {
	return len(b)
}

// Swap is part of the sorting algorithm.
func (b ByVersion) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less is part of the sorting algorithm.
func (b ByVersion) Less(i, j int) bool {
	cmp := version.CompareSimple(version.Normalize(b[i].ID), version.Normalize(b[j].ID))

	if cmp == 0 {
		return b[i].ID < b[j].ID
	}

	return cmp < 0
}

// ByMinecraft sorts a list of Forge versions by ID.
type ByMinecraft Versions

// Sort simply sorts the versions list.
func (b ByMinecraft) Sort() {
	sort.Sort(b)
}

// Len is part of the sorting algorithm.
func (b ByMinecraft) Len() int {
	return len(b)
}

// Swap is part of the sorting algorithm.
func (b ByMinecraft) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less is part of the sorting algorithm.
func (b ByMinecraft) Less(i, j int) bool {
	cmp := version.CompareSimple(version.Normalize(b[i].Minecraft), version.Normalize(b[j].Minecraft))

	if cmp == 0 {
		return b[i].Minecraft < b[j].Minecraft
	}

	return cmp < 0
}
