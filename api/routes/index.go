package routes

import "github.com/jpenaroche/go-simple-api/utils"

var Routes map[string][]utils.Route = map[string][]utils.Route{
	"persons": Persons,
}
