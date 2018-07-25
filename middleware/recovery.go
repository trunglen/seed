package middleware

import (
	"rocky-springs-86767/g/x/web"
	"rocky-springs-86767/x/logger"

	"github.com/gin-gonic/gin"
)

var panicLog = logger.NewLogger("panic")

func RecoveryWithWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if err, same := err.(error); same {
					if webError, same := err.(web.IWebError); same {
						c.AbortWithStatusJSON(webError.StatusCode(), map[string]interface{}{
							"error":  err.Error(),
							"status": "error",
						})
					} else {
						c.AbortWithStatusJSON(500, map[string]interface{}{
							"error":  err.Error(),
							"status": "error",
						})
					}
				} else {
					panicLog.Error(err)
				}
			}
		}()
		c.Next()
	}
}
