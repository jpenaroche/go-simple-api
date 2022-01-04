package routes

import "github.com/jpenaroche/go-simple-api/utils"

var Routes map[string][]utils.RouteParameter = map[string][]utils.RouteParameter{
	"persons": Persons,
}
