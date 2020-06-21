package main

import (
	"net"
	"net/http"
	"os"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {

	baseTmpl := "templates/base.tmpl.html"
	homeTmpl := "templates/home.tmpl.html"
	aboutTmpl := "templates/about.tmpl.html"
	nameTmpl := "templates/name.tmpl.html"
	s404Tmpl := "templates/404.tmpl.html"

	r := multitemplate.NewRenderer()
	r.AddFromFiles("home", baseTmpl, homeTmpl)
	r.AddFromFiles("404", baseTmpl, s404Tmpl)
	r.AddFromFiles("about", baseTmpl, aboutTmpl)
	r.AddFromFiles("name", baseTmpl, nameTmpl)
	return r
}

func main() {
	router := gin.Default()

	router.Use(gin.Logger())

	//router.LoadHTMLGlob("./templates/*.tmpl.html")
	router.Static("/static", "./static")

	router.HTMLRender = createMyRender()

	//initializeRoutes()
	router.GET("/", func(ctx *gin.Context) {
		hostname, _ := os.Hostname()
		ip, _ := net.LookupIP(hostname)
		ctx.HTML(
			http.StatusOK,
			"home",
			gin.H{
				"title":    "Home",
				"hostname": hostname,
				"IP":       ip,
			},
		)

	})

	router.GET("/about", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"about",
			gin.H{
				"title": "About",
			},
		)

	})

	// Test Json Get Call
	router.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"message": "OK",
			})
	})

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.HTML(
			http.StatusOK,
			"name",
			gin.H{
				"title": "user",
				"name":  name,
			},
		)

	})

	//health route for kubernetes
	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(
			http.StatusOK,
			"Working!")
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(
			404,
			"404",
			gin.H{
				"title": "404",
			},
		)
	})

	router.Run()

}
