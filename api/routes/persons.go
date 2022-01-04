package routes

import (
	"github.com/jpenaroche/go-simple-api/api/handlers"
	"github.com/jpenaroche/go-simple-api/utils"
)

var Persons []utils.RouteParameter = []utils.RouteParameter{
	{
		Path:    `/persons/{id:\d+}`,
		Verb:    utils.Get,
		Handler: handlers.GetPerson,
	},
	{
		Path:    `/people/{id:\d+}`,
		Verb:    utils.Get,
		Handler: handlers.GetPerson,
	},
}
