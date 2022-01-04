package config

import (
	"github.com/joho/godotenv"
)

type EnvVar string

type Config struct {
	Server *Server
}

func (c *Config) Load() *Config {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	c.Server = new(Server).Load() //Injecting Server config
	return c
}
