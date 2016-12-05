package handling

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"github.com/lcaballero/statics/files"
	"path"
)

const DefaultIndexName = "index.html"
const IndexPage = "/index.html"

// RouteFileSystem is a http.FileSystem using the given route
// directory to find files.
type RouteFileSystem struct {
	local string
	remote string
	Index string
}

// NewRouteFileSystem maps calls through a route to the local file
// system.  The parameters local refers to the local directory where
// assets reside, and remote refers to where the client believes
// assets live.  It basically remaps remote directory to the local
// equivalent.
func NewRouteFileSystem(local, remote string) *RouteFileSystem {
	pre := strings.TrimSuffix(remote, "/")

	//TODO: make sure local has form '/dir' with '/' suffix
	//TODO: make sure remote has form '/dir/' with '/' prefix and suffix
	rf := &RouteFileSystem{
		local: local,
		remote: pre,
		Index: DefaultIndexName,
	}
	return rf
}

// IndexName provides the name of the index file if one has been
// set and only if the requested name is not provided, as in
// the special case for '/'.
func (sf *RouteFileSystem) IndexName() string {
	if sf.Index == "" {
		return DefaultIndexName
	} else {
		return sf.Index
	}
}

// Open turns the FileProvider interface into a http.FileSystem.
func (sf *RouteFileSystem) Open(name string) (http.File, error) {
	if !strings.HasPrefix(name, sf.remote) {
		return nil, fmt.Errorf("request didn't have prefix: %s", sf.remote)
	}
	// Have: some name like '/<prefix>/file.ext'

	fmt.Println("open name:", name, sf.local, sf.remote)

	if name == sf.remote {
		name = sf.IndexName()
	} else {
		name = name[len(sf.remote):]
		// Have: file like '/file.ext'
		name = strings.TrimPrefix(name, "/")
		// Have: file.ext
	}

	file := filepath.Join(sf.local, name)
	fmt.Println("attempting to open:", file)
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	return f, err
}

func (sf *RouteFileSystem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	files.NewFileHandler(w, r, sf.Index).
		ServeFile(sf, path.Clean(upath), true)
}
