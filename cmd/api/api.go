package main

import (
	"flag"
	"log"
	"notegin/internal/routes"
)

func main() {
	port := flag.String("port", "8001", "you can set new port into server")
	log.Println("server has started on localhost:" + *port)

	// if err := godotenv.Load("./internal/config/.env"); err != nil {
	// 	log.Fatal(err)
	// }

	router := routes.Init()

	if err := router.Run("localhost:" + *port); err != nil {
		log.Fatal(err)
	}

}
