package handling

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// BasicStaticRoute sets up a route for serving static assets
// directly from the file system.
func BasicStaticRoute(root, prefix, index string, route *mux.Route) {
	r := route.PathPrefix(prefix).
		HandlerFunc(FromRoot(root, prefix)).
		Methods("GET")

	err := r.GetError()
	if err != nil {
		panic(err)
	}
}

// FromRoot creates a handler with a file system rooted at the
// provided root location.
func FromRoot(local, remote string) http.HandlerFunc {
	fs := NewRouteFileSystem(local, remote)
	return HandleStatics(local, fs)
}

// HandleStatics internally creates a FileServer to handle serving
// static assets.  However, requests are first parsed given the
// routing parameters and a new file location is used to actually
// find the file on disc.  This Handler requires that the route
// keys 'hash', 'name' and 'ext' can be found in the request
// variables.
func HandleStatics(root string, fs http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("\nrequested: ", req.RequestURI)

		vars := Vars{
			AssetVars: AssetVars{},
			req:       req,
		}

		// By definition Parts has to include at least [assets] as prefix.
		parts := Path(req.URL.Path).Parts()
		vars.AcquireVars(parts)

		if vars.IsDebugOn() {
			dbg := debug{res: res, req: req}
			dbg.ToLog()
		}

		fmt.Println(parts, req.URL.Path, req.RequestURI)
		filereq := vars.RewritePath(req)
		fmt.Println("serving file:", filereq.URL.Path, filereq.RequestURI)
		fs.ServeHTTP(res, filereq)
	}
}
