package middleware

import (
	"awesomeProject/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.JSON(appErr.StatusCode, appErr)
					return
				}

				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": http.StatusInternalServerError,
					"message":     "Internal server error",
					"log":         err,
				})
			}
		}()
	}
}
