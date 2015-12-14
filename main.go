package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/nabeken/negroni-auth"
	"github.com/pivotal-pez/haas-broker/handlers/binding"
	"github.com/pivotal-pez/haas-broker/handlers/catalog"
	"github.com/pivotal-pez/haas-broker/handlers/instance"
)

func main() {
	router := mux.NewRouter()
	n := negroni.Classic()

	router.HandleFunc(catalog.HandlerPath, catalog.Get()).Methods("GET")
	router.HandleFunc(instance.AsyncHandlerPath, instance.Get()).Methods("GET")
	router.HandleFunc(instance.HandlerPath, instance.Put()).Methods("PUT")
	router.HandleFunc(instance.HandlerPath, instance.Patch()).Methods("PATCH")
	router.HandleFunc(instance.HandlerPath, instance.Delete()).Methods("DELETE")
	router.HandleFunc(binding.HandlerPath, binding.Delete()).Methods("DELETE")
	router.HandleFunc(binding.HandlerPath, binding.Put()).Methods("PUT")

	n.Use(auth.Basic("username", "secretpassword"))
	n.UseHandler(router)
	n.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
