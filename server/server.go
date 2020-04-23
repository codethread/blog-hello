package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const assetDir = "/assets"

// Create a Server
func Create() *gin.Engine {
	return gin.Default()
}

// SetupFileHandler adds static file handling to server
func SetupFileHandler(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.Static(assetDir, "./assets")
}

// SetupRoutes mount routes on an server
func SetupRoutes(r *gin.Engine) {
	r.GET("/", viewHome)
	r.GET("/api/__healthcheck", healthCheck)
}

func viewHome(c *gin.Context) {
	if pusher := c.Writer.Pusher(); pusher != nil {
		files := []string{"app.js", "main.css"}
		for _, f := range files {
			path := fmt.Sprintf("%s/%s", assetDir, f)
			if err := pusher.Push(path, nil); err != nil {
				fmt.Printf("Failed to push: %v", err)
			}
		}
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Blog",
	})
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": "Ok",
	})
}
