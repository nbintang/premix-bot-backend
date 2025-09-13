package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// TraceIDMiddleware adds a trace_id for each request and logs it using Logrus
func TraceIDMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		traceID := c.GetHeader("X-Trace-Id")
		if traceID == "" {
			traceID = uuid.New().String()
		}
		c.Set("trace_id", traceID)
		c.Writer.Header().Set("X-Trace-Id", traceID)

		logger.WithFields(logrus.Fields{
			"trace_id":  traceID,
			"method":    c.Request.Method,
			"path":      c.Request.URL.Path,
			"client_ip": c.ClientIP(),
		}).Info("incoming request")

		c.Next()

		// Log response
		status := c.Writer.Status()
		logger.WithFields(logrus.Fields{
			"trace_id": traceID,
			"status":   status,
		}).Info("request completed")
	}
}
