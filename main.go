package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-github-actions-example/handlers"
	"github.com/shaikrasheed99/golang-github-actions-example/routes"
)

func main() {
	app := gin.Default()

	th := handlers.NewTestHandler()

	routes.RegisterTestRoutes(app, th)

	app.Run()
}
