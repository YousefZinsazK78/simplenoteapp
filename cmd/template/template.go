package main

import (
	"flag"
	"log"
	"notegin/internal/routes/tmplate"
)

func main() {
	port := flag.String("port", "8000", "you can set new port into server")
	log.Println("server has started on localhost:" + *port)

	router := tmplate.Init()

	if err := router.Run("localhost:" + *port); err != nil {
		log.Fatal(err)
	}
}
