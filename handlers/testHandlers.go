package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-github-actions-example/constants"
)

type TestHandler interface {
	FirstTestHandler(*gin.Context)
}

type testHandler struct{}

func NewTestHandler() TestHandler {
	return &testHandler{}
}

func (th *testHandler) FirstTestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, constants.FisrtTestAPIMessage)
}