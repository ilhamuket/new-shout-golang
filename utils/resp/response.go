package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseFormatter(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func Success(ctx *gin.Context, data interface{}) {
	ResponseFormatter(ctx, http.StatusOK, "Success", data)
}

func Error(ctx *gin.Context, status int, message string) {
	ResponseFormatter(ctx, status, message, nil)
}
