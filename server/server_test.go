package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthRoute(t *testing.T) {
	r := gin.Default()
	SetupRoutes(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/__healthcheck", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"health\":\"Ok\"}", w.Body.String())
}
