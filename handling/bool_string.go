package handling

import (
	"strings"
)

// BoolString is intended to map: 'on', '1', 'yes', 'y', 'ok' to true.
type BoolString string

// IsTrue returns true if the BoolString is one of 'on', '1', 'yes', 'y', 'ok',
// else it returns false.
func (b BoolString) IsTrue() bool {
	switch strings.ToLower(string(b)) {
	case "on", "1", "yes", "y", "ok":
		return true
	default:
		return false
	}
}
