package main

import (
	"github.com/gorilla/mux"
	"go-admin/middleware/routerBoot"
	_ "go-admin/router/auth"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	routerBoot.RegisterEngine(router)
	routerBoot.StartRoute()
	log.Fatal(http.ListenAndServe(":8080", router))
}
