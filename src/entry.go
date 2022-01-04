package main

import (
	"github.com/jpenaroche/go-simple-api/src/config"
)

type Context struct {
	Config *config.Config
}

func main() {
	config := new(config.Config).Load()

	ctx := Context{
		Config: config,
	}
	Run(ctx) //TODO inject Context
}
