package persons

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jpenaroche/go-simple-api/api/schemas"
	"github.com/jpenaroche/go-simple-api/utils"
)

var people []schemas.Person = []schemas.Person{
	{
		Id:   1,
		Name: "John",
		Age:  30,
	},
	{
		Id:   2,
		Name: "Jane",
		Age:  25,
	},
}

func GetPerson(w http.ResponseWriter, r *utils.Request) {
	fmt.Println("GetPerson id:", r.Params["id"])
	for _, person := range people {
		if strconv.Itoa(person.Id) == r.Params["id"] {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(person)
			return
		}
	}

	utils.GetErrorResponse(w, "Person Not Fond", http.StatusNotFound)
}
