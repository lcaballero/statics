package handling

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
)

// RouteFileSystem is a http.FileSystem using the given route
// directory to find files.
type RouteFileSystem struct {
	Root string
}

// Open turns the FileProvider interface into a http.FileSystem.
func (sf RouteFileSystem) Open(name string) (http.File, error) {
	file := filepath.Join(sf.Root, name)
	fmt.Println("attempting to open: ", name, file)
	if name == "/" {

	}
	return os.Open(file)
}

// BasicStaticRoute sets up a route for serving static assets
// directly from the file system.
func BasicStaticRoute(root string, route *mux.Route) {
	prefix := "/assets/"
	r := route.PathPrefix(prefix).
		MatcherFunc(AssetMatcher(prefix)).
		HandlerFunc(FromRoot(root)).
		Methods("GET")

	err := r.GetError()
	if err != nil {
		panic(err)
	}
}

// FromRoot creates a handler with a file system rooted at the
// provided root location.
func FromRoot(root string) http.HandlerFunc {
	fs := RouteFileSystem{Root: root}
	fserve := http.FileServer(fs)
	return HandleStatics(root, fserve)
}

// HandleStatics internally creates a FileServer to handle serving
// static assets.  However, requests are first parsed given the
// routing parameters and a new file location is used to actually
// find the file on disc.  This Handler requires that the route
// keys 'hash', 'name' and 'ext' can be found in the request
// variables.
func HandleStatics(root string, fs http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var vars Vars = Vars{
			vars: mux.Vars(req),
			req:  req,
		}
		parts := Path(req.URL.Path).Parts()[1:]
		vars.vars.AcquireVars(parts)

		if vars.IsDebugOn() {
			dbg := debug{res: res, req: req}
			dbg.ToLog()
		}

		filereq := vars.RewritePath(req)
		fmt.Println("vars:", vars.vars, filereq.URL.Path, filereq.RequestURI)
		fs.ServeHTTP(res, filereq)
		fmt.Println("should served: ", root, filereq.URL.Path, filereq.RequestURI)
	}
}
