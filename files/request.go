package files

import "net/http"


type Request struct {
	*http.Request
}

func (r Request) Val(key string) string {
	if v := r.Request.Header[key]; len(v) > 0 {
		return v[0]
	}
	return ""
}

