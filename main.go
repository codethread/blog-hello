package main

import (
	"fmt"
	"github.com/AHDesigns/blogui/server"
)

func main() {
	r := server.Create()
	server.SetupFileHandler(r)
	server.SetupRoutes(r)
	if err := r.RunTLS(
		":443",
		"./certs/localhost.pem",
		"./certs/localhost-key.pem",
	); err != nil {
		fmt.Printf("Failed to start server: err\n %s", err)

	}
}
