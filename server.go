package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jpenaroche/go-simple-api/utils"
)

const port = "8080"
const server = "localhost"

func Run(modules map[string][]utils.Route) {
	for module, routes := range modules {
		log.Println("Routing module", module)
		for _, route := range routes {
			http.HandleFunc("/", utils.Bypass(route))
		}
	}
	_ = http.ListenAndServe(fmt.Sprintf("%s:%s", server, port), nil)
}
