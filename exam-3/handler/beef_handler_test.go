package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/beef/summary", GetBeefSummary)
	return r
}

func TestGetBeefSummary(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/beef/summary", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "beef")

	body := w.Body.String()
	assert.True(t, strings.Contains(body, `"t-bone":`) || strings.Contains(body, `"T-bone":`), "should contain t-bone (case insensitive)")
	assert.Contains(t, body, "fatback")
	assert.Contains(t, body, "pastrami")
	assert.Contains(t, body, "pork")
	assert.Contains(t, body, "meatloaf")
	assert.Contains(t, body, "jowl")
	assert.Contains(t, body, "enim")
	assert.Contains(t, body, "bresaola")
}
