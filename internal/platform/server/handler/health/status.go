package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StatusHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Everything is OK!")
	}
}
