package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestIDKey is key for request-id middleware
const RequestIDKey = "X-Request-ID"

// MakeRequestIDGinMiddleware initializes the RequestID middleware
func MakeRequestIDGinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rid string
		passedRequestID := c.GetHeader(RequestIDKey)
		if len(passedRequestID) > 0 {
			rid = passedRequestID
		} else {
			rid = uuid.New().String()
		}

		c.Writer.Header().Set(RequestIDKey, rid)
		c.Request.Header.Set(RequestIDKey, rid)
		c.Set(RequestIDKey, rid)

		c.Next()
	}
}
