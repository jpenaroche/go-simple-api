package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jpenaroche/go-simple-api/api/routes"
	"github.com/jpenaroche/go-simple-api/utils"
)

const port = "8080"
const server = "localhost"

func Run() {
	handler := new(utils.Router).ProcessRoutes(routes.Routes)
	log.Println("Server running on http://" + server + ":" + port)
	_ = http.ListenAndServe(fmt.Sprintf("%s:%s", server, port), handler)
}
