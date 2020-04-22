package main

import "github.com/AHDesigns/blogui/server"

func main() {
	r := server.Create()
	server.SetupFileHandler(r)
	server.SetupRoutes(r)
	r.RunTLS(
		":8080",
		"./certs/localhost.pem",
		"./certs/localhost-key.pem",
	)
}
