package handling

import (
	"path/filepath"
	"strings"
)

// Path represents a file system path.
type Path string

// Split delegates to filepath.Split for the directory and file.
func (p Path) Split() (string, string) {
	return filepath.Split(string(p))
}

// Dirs creates a Parts collection of dirs in the Path.
func (p Path) Dirs() (Parts, string) {
	dir, file := p.Split()
	return strings.Split(dir, "/"), file
}

// Parts provides the Path dirs as a Parts type.
func (p Path) Parts() Parts {
	pcs, file := p.Dirs()
	parts := []string{}
	for _, p := range pcs {
		if p != "" {
			parts = append(parts, p)
		}
	}
	if file != "" {
		parts = append(parts, file)
	}
	return parts
}
