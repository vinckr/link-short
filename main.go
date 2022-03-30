package main

import (
	"log"

	"github.com/vinckr/link-short/api"
	"github.com/vinckr/link-short/config"
)

func main() {
	c, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}
	api, err := api.New(c)
	if err != nil {
		log.Fatal(err)
	}
	api.Start()
}
