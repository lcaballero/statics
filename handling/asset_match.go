package handling

import (
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
	"fmt"
)

// AssertMatcher creates a Matcher that will match the given prefix
// but additionally add route vars intended for serving static assets.
func AssetMatcher(prefix string) mux.MatcherFunc {
	return func(req *http.Request, m *mux.RouteMatch) bool {
		vars := AssetVars{}
		path := req.URL.Path

		isMatch := vars.HasPrefix(path, prefix)
		if !isMatch {
			return false
		}

		p := filepath.Clean(path)
		parts := Path(p).Parts()

		fmt.Println("asset matcher", req.URL.Path, path, p, parts)

		return true
	}
}

