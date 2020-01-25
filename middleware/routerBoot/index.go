package routerBoot

import "github.com/gorilla/mux"

var routeList []func(router *mux.Router)
var bootEngine *mux.Router

func RegisterEngine(engine *mux.Router) {
	bootEngine = engine
}

func InjectRoute(callback func(router *mux.Router)) {
	routeList = append(routeList, callback)
}

func StartRoute() {
	for _, cb := range routeList {
		cb(bootEngine)
	}
}
