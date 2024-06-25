package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"order/controllers"
	"order/mappings"
	"order/model"
	"strings"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := controllers.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

// Test for POST /order/add
func TestPostUser(t *testing.T) {
	mappings.CreateUrlMappings()
	router := mappings.Router
	w := httptest.NewRecorder()

	// Create an example order for testing
	exampleUser := model.OrderReq{
		OrderType: "1",
		Symbol:    "BTCUSDT",
	}
	userJson, _ := json.Marshal(exampleUser)
	req, _ := http.NewRequest("POST", "/v1/order/add", strings.NewReader(string(userJson)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
