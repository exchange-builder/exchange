package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"order/model"
	"order/web"
	"strings"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := web.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

// Test for POST /order/add
func TestPostUser(t *testing.T) {
	router := web.SetupRouter()
	router = web.PostOrder(router)

	w := httptest.NewRecorder()

	// Create an example order for testing
	exampleUser := model.OrderReq{
		OrderType: "1",
		Symbol:    "BTCUSDT",
	}
	userJson, _ := json.Marshal(exampleUser)
	req, _ := http.NewRequest("POST", "/order/add", strings.NewReader(string(userJson)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
