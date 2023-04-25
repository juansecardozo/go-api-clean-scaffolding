package main

import (
	"github.com/juansecardozo/go-api-clean-scaffolding/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
