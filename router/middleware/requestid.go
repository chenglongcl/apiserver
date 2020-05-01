package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Check for incoming header,If X-request-id exists
		requestId := c.Request.Header.Get("X-Request-ID")
		if requestId == "" {
			u4 := uuid.NewV4()
			requestId = u4.String()
		}

		// Expose it for use in the application
		c.Set("X-Request-ID", requestId)
		// Set X-Request-ID Header
		c.Writer.Header().Set("X-Request-ID", requestId)
		c.Next()
	}
}
