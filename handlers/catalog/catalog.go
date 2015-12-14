package catalog

import (
	"fmt"
	"net/http"
)

const (
	HandlerPath = "/v2/catalog"
)

func Get() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	}
}
