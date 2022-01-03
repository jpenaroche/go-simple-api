package routes

import (
	"github.com/jpenaroche/go-simple-api/api/handlers/persons"
	"github.com/jpenaroche/go-simple-api/utils"
)

var Persons []utils.Route = []utils.Route{
	{
		Path:    `/persons/{id:\d+}`,
		Verb:    utils.Get,
		Handler: persons.GetPerson,
	},
}
