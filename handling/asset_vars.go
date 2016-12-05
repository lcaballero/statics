package handling

import (
	"fmt"
	"strings"
)

// AssetVars key/value pairs.
type AssetVars map[string]string

// get retrieves the key from the map and returns it if it's
// present else an empty string.
func (a AssetVars) get(key string) string {
	val, ok := a[key]
	if ok {
		return val
	} else {
		return ""
	}
}

// File returns the value of 'file'
func (a AssetVars) File() string {
	return a.get("file")
}

// Path returns the value of 'path'
func (a AssetVars) Path() string {
	return a.get("path")
}

func (m AssetVars) AcquireVars(parts Parts) {
	fmt.Println("acquire vars:", parts, parts.Len())

	m["path"] = "/" + strings.Join(parts.NighAll(), "/")
	m["file"] = parts.Last()
}
