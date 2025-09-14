package utils

import "github.com/gin-gonic/gin"

// ResponseHelper helper struct
type ResponseHelper struct{}

// NewResponseHelper init
func NewResponseHelper() *ResponseHelper {
	return &ResponseHelper{}
}

// SuccessResponse returns a success response
func (r *ResponseHelper) SuccessResponse(c *gin.Context, statusCode int, msg string, data any) {
	c.JSON(statusCode, gin.H{
		"success": true,
		"message": msg,
		"data":    data,
	})
}

// SuccessWithMeta returns a success response with meta data
func (r *ResponseHelper) SuccessWithMeta(c *gin.Context, statusCode int, msg string, data any, meta any) {
	c.JSON(statusCode, gin.H{
		"success": true,
		"message": msg,
		"data":    data,
		"meta":    meta,
	})
}

// ErrorResponse returns an error response
func (r *ResponseHelper) ErrorResponse(c *gin.Context, statusCode int, msg string, err any) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"message": msg,
		"error":   err,
	})
}
