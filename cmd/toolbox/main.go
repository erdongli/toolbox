package main

import (
	"log"

	"github.com/erdongli/toolbox/internal/server"
)

func main() {
	log.Fatal(server.Serve())
}
