package handling

import (
	"fmt"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
)

type Debug interface {
	ToLog()
	IsDebugOn() bool
}

type debug struct {
	res http.ResponseWriter
	req *http.Request
}

func (d debug) ToLog() {
	req := d.req
	vars := mux.Vars(req)

	for k, v := range vars {
		fmt.Println(k, ": ", v)
	}

	ext, _ := vars["ext"]
	mime := mime.TypeByExtension("." + ext)

	fmt.Printf("ext: %s\n", ext)
	fmt.Printf("mime-type: %s\n", mime)
}
