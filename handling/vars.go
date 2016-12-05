package handling

import (
	"fmt"
	"net/http"
	"path/filepath"
)

// DefaultDebugKey the query parameter key name used by default but can be
// overwritten in a vars instance.
const DefaultDebugKey = "statics_debug"

// Vars represent a wrapper around http request parameters that
// can be used to make decisions about how to serve static files.
type Vars struct {
	AssetVars
	req      *http.Request
	debugKey string
}

// IsDebugOn checks for the parameter 'debug' in the query string, if
// it finds the key with a value of 'on', '1', 'yes', 'y', 'ok' it
// considers debug to be on and return true; false otherwise.
func (v Vars) IsDebugOn() bool {
	if v.req == nil || v.req.URL == nil {
		return false
	}

	vars := v.req.URL.Query()
	if vars == nil {
		return false
	}

	key := v.debugKey
	if key == "" {
		key = DefaultDebugKey
	}

	db, ok := vars[key]
	if !ok {
		return false
	}

	if len(db) != 1 {
		return false
	}

	return BoolString(db[0]).IsTrue()
}

// RewritePath uses 'ext' and 'name' to create a new request looking
// specifically for a file built from those values instead of using
// the raw query string.
func (v Vars) RewritePath(req *http.Request) *http.Request {
	//TODO: return original req if 'ext' key doesn't exist
	path := v.AssetVars["path"]

	//TODO: return original req if 'name' key doesn't exist
	file := v.AssetVars["file"]

	newfile := filepath.Join(path, file)
	fmt.Println("rewrite path:", req.RequestURI, req.URL.Path, newfile)

	req2 := *req
	url2 := *req.URL
	url2.Path = newfile
	req2.RequestURI = newfile
	req2.URL = &url2

	return &req2
}
