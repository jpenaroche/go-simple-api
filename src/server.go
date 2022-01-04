package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jpenaroche/go-simple-api/src/app/routes"
	"github.com/jpenaroche/go-simple-api/src/utils"
)

func Run(ctx Context) {
	host := ctx.Config.Server.Host
	port := ctx.Config.Server.Port
	handler := new(utils.Router).ProcessRoutes(routes.Routes)
	log.Println("Server running on http://" + host + ":" + port)
	_ = http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), handler)
}
