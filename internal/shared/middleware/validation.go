package middleware

import (
	"net/http"
	"premix-backend/internal/shared/utils"
	"reflect"

	"github.com/gin-gonic/gin"
)

// Validation struct
type Validation struct {
	ResponseHelper *utils.ResponseHelper
}

// NewValidation constructor
func NewValidation(responseHelper *utils.ResponseHelper) *Validation {
	return &Validation{ResponseHelper: responseHelper}
}

// Middleware function for validation
func (v *Validation) Middleware(obj any) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := reflect.TypeOf(obj).Elem()
		newObj := reflect.New(t).Interface()

		if err := c.ShouldBind(newObj); err != nil {
			v.ResponseHelper.ErrorResponse(c, http.StatusBadRequest, "Validation error", err.Error())
			c.Abort()
			return
		}

		c.Set("body", newObj)
		c.Next()
	}
}
