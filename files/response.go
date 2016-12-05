package files

import (
	"net/http"
	"fmt"
	"net/url"
	"sort"
)


type Response struct {
	http.ResponseWriter
}

func (r Response) Val(key string) string {
	h := r.ResponseWriter.Header()
	if h == nil {
		return ""
	}
	if v := h[key]; len(v) > 0 {
		return v[0]
	}
	return ""
}

// Error replies to the request with the specified error message and HTTP code.
// The error message should be plain text.
func (r Response) Error(error string, code int) {
	w := r.ResponseWriter
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, error)
}

func (w Response) DirList(f http.File) {
	dirs, err := f.Readdir(-1)
	if err != nil {
		// TODO: log err.Error() to the Server.ErrorLog, once it's possible
		// for a handler to get at its Server via the ResponseWriter. See
		// Issue 12438.
		w.Error("Error reading directory", http.StatusInternalServerError)
		return
	}
	sort.Sort(ByName(dirs))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<pre>\n")
	for _, d := range dirs {
		name := d.Name()
		if d.IsDir() {
			name += "/"
		}
		// name may contain '?' or '#', which must be escaped to remain
		// part of the URL path, and not indicate the start of a query
		// string or fragment.
		url := url.URL{Path: name}
		fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", url.String(), htmlReplacer.Replace(name))
	}
	fmt.Fprintf(w, "</pre>\n")
}
