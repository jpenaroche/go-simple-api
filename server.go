package main

import (
	"fmt"
	"net/http"

	"github.com/jpenaroche/go-simple-api/utils"
)

const port = "8080"
const server = "localhost"

func Run(modules map[string][]utils.RouteParameter) {
	handler := new(utils.Router).ProcessRoutes(modules)
	_ = http.ListenAndServe(fmt.Sprintf("%s:%s", server, port), handler)
}
