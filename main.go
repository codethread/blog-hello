package main

import (
	"fmt"
	"github.com/AHDesigns/blogui/server"
)

func main() {
	r := server.Create()
	server.SetupFileHandler(r)
	server.SetupRoutes(r)
	r.RunTLS(
		":443",
		"./certs/localhost.pem",
		"./certs/localhost-key.pem",
	)

	// if err := r.RunTLS(
	// 	":443",
	// 	"./certs/localhost.pem",
	// 	"./certs/localhost-key.pem",
	// ); err != nil {
	// 	fmt.Printf("failed to start server: %s", err)

	// }
}
