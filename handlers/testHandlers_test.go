package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/golang-github-actions-example/constants"
	"github.com/stretchr/testify/assert"
)

func TestFirstTestHandler(t *testing.T) {
	t.Run("should be able return first test api with success message", func(t *testing.T) {
		th := NewTestHandler()
		router := gin.New()
		router.GET(constants.FirstTestAPIPath, th.FirstTestHandler)

		req, _ := http.NewRequest("GET", constants.FirstTestAPIPath, nil)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		var resMessage string
		_ = json.Unmarshal(res.Body.Bytes(), &resMessage)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, constants.FisrtTestAPIMessage, resMessage)
	})
}

func TestSecondTestHandler(t *testing.T) {
	t.Run("should be able return second test api with success message", func(t *testing.T) {
		th := NewTestHandler()
		router := gin.New()
		router.GET(constants.SecondTestAPIPath, th.SecondTestHandler)

		req, _ := http.NewRequest("GET", constants.SecondTestAPIPath, nil)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		var resMessage string
		_ = json.Unmarshal(res.Body.Bytes(), &resMessage)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, constants.SecondTestAPIMessage, resMessage)
	})
}
