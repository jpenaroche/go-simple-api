package schemas

import (
	"log"
)

func (p *Person) tellMyName() string {
	return p.Name
}

func (p *Person) doNothing() {
	log.Print("do nothing")
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
