package binding

import (
	"fmt"
	"net/http"
)

const (
	//HandlerPath - path to normal instance handlers
	HandlerPath = "/v2/service_instances/{instance_id}/service_bindings/{binding_id}"
)

//Put - handler function for put calls
func Put() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	}
}

//Delete - handler function for delete calls
func Delete() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	}
}
