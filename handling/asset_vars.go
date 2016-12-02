package handling
import (
	"strings"
	"path/filepath"
	"fmt"
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
	if len(parts) == 1 {
		m["path"] = "/"
		m["file"] = parts.Last()
		return
	}
	if len(parts) > 1 {
		m["path"] = "/" + strings.Join(parts.NighAll(), "/")
		m["file"] = parts.Last()
		return
	}
}

func (m AssetVars) HasPrefix(path, prefix string) bool {
	p := filepath.Clean(path)
	pre := filepath.Clean(prefix)
	fmt.Println("path", path, "p", p, "pre", pre)
	return strings.HasPrefix(p, pre)
}
