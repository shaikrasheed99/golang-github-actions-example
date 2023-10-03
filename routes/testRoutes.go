package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-github-actions-example/constants"
	"github.com/shaikrasheed99/golang-github-actions-example/handlers"
)

func RegisterTestRoutes(engine *gin.Engine, th handlers.TestHandler) {
	engine.GET(constants.FirstTestAPIPath, th.FirstTestHandler)
	engine.GET(constants.SecondTestAPIPath, th.SecondTestHandler)
	engine.GET(constants.ThirdTestAPIPath, th.ThirdTestHandler)
}
