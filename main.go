package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin to production mode
	// gin.SetMode(gin.ReleaseMode)
	// Get the gin router
	r := gin.Default()
	// add gzip to load files faster
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	// parse and execute templates
	r.LoadHTMLGlob("production/*.gohtml")
	// grab static components
	r.Static("/dist", "./dist")
	// control home route
	r.GET("/", Index)
	// Handle all other routes
	r.NoRoute(NotFound)

	// To run locally uncomment the line below and change init func to main
	r.Run(":3000")

	// To work with App Engine
	http.Handle("/", r)

}

// Index loads the site
func Index(c *gin.Context) {
	c.HTML(200, "resume.gohtml", gin.H{
		"title": "Deliri Resume",
	})
}

// NotFound handles improper route requests
func NotFound(c *gin.Context) {
	c.HTML(404, "404.gohtml", gin.H{
		"title": "404",
	})
	return
}
