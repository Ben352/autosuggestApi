package handler

import (
	"fmt"
	"net/http"

	"github.com/Ben352/autosuggestApi/internal/autocomplete"
	"github.com/Ben352/autosuggestApi/internal/middleware"
	"github.com/gin-gonic/gin"
)

type AddWordBody struct {
	WordToAdd string `json:"wordToAdd" binding:"required"`
}
type loadFileBody struct {
	FileName string `json:"fileName" binding:"required"`
}

func RegisterRoutes(r *gin.Engine, apiKey string) {

	apiMiddleware := middleware.ValidateAPIKeyMiddleware(apiKey)
	r.GET("/autocomplete", autocompleteHandler)

	// admin rotues
	admin := r.Group("/admin", apiMiddleware)
	admin.POST("/addWord", addWordHandler)
	admin.POST("/save", saveTrieHandler)
	admin.POST("/load", loadTrieHandler)
}

func saveTrieHandler(c *gin.Context) {
	var req loadFileBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	wordToAdd := req.FileName
	autocomplete.SaveTrie(wordToAdd)
	c.JSON(http.StatusOK, gin.H{"message": "Trie saved successfully"})

}

func loadTrieHandler(c *gin.Context) {

	var req loadFileBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	wordToAdd := req.FileName
	autocomplete.LoadTrie(wordToAdd)
	c.JSON(http.StatusOK, gin.H{"message": "word added successfully"})
}

func addWordHandler(c *gin.Context) {
	var req AddWordBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	wordToAdd := req.WordToAdd
	autocomplete.AddWord(wordToAdd)
	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("added %s", wordToAdd)})
}

func autocompleteHandler(c *gin.Context) {
	prefix := c.Query("prefix")
	suggestions := autocomplete.GetSuggestions(prefix)
	c.JSON(200, gin.H{"suggestions": suggestions})
}
