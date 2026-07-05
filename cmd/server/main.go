package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"time":   time.Now().UTC().Format(time.RFC3339),
		})
	})
	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
