package auth

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-admin/middleware/routerBoot"
	"net/http"
)

type LoginOutController struct{}

func init() {
	routerBoot.InjectRoute(new(LoginOutController).registerRoute)
}

func (controller *LoginOutController) registerRoute(router *mux.Router) {
	router.HandleFunc("/login", index).Methods("get")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}
