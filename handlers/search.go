package handlers

import (
	"infra-search/search"
	"net/http"

	"github.com/gin-gonic/gin"
)

type searchRequest struct {
	Query string `json:"query" binding:"required"`
}

func SearchHandler(c *gin.Context) {
	var req searchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query は必須です"})
		return
	}

	query := search.BuildQuery(req.Query)
	results, err := search.Search(query)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"results": []interface{}{},
			"message": err.Error(),
		})
		return
	}

	if len(results) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"results": []interface{}{},
			"message": "結果が見つかりませんでした",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}
