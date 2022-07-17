package main

import (
	"log"

	"github.com/fransafu/golang-dragonflydb-example/internal/app"
)

func main() {
	if err := app.StartApp(); err != nil {
		log.Fatal(err)
	}
}
