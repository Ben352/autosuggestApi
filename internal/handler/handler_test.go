package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	RegisterRoutes(r, "test-Api-key")
	return r
}

func TestAddWordHandler(t *testing.T) {
	r := setupRouter()

	// Test adding a word
	word := AddWordBody{WordToAdd: "Hello Test"}
	jsonValue, _ := json.Marshal(word)
	req, _ := http.NewRequest("POST", "/admin/addWord", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", "test-Api-key")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	fmt.Println(response)
	assert.Equal(t, fmt.Sprintf("added %s", word.WordToAdd), response["message"])
}
