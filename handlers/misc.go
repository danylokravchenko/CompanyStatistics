package handlers

import (
	"github.com/UndeadBigUnicorn/CompanyStatistics/cache"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	appCache *cache.Cache
)

func init() {
	appCache = cache.New()
}


// 400 wrapper
func _400(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "400 Bad Request",
	})
}


// 404 wrapper
func _404(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": message,
	})
}