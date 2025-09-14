package auth

import (
	"net/http"
	"premix-backend/internal/auth/dto"
	"premix-backend/internal/shared/utils"

	"github.com/gin-gonic/gin"
)

// HandlerImpl struct
type HandlerImpl struct {
	service        Service
	responseHelper *utils.ResponseHelper
}

// NewHandler init
func NewHandler(service Service, responseHelper *utils.ResponseHelper) *HandlerImpl {
	return &HandlerImpl{service: service, responseHelper: responseHelper}
}

// Login for login
func (h *HandlerImpl) Login(c *gin.Context) {
	req, _ := c.Get("body")
	loginReq := req.(*dto.LoginRequest)

	token, err := h.service.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		h.responseHelper.ErrorResponse(c, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	h.responseHelper.SuccessResponse(c, http.StatusOK, "Success Login", token)

}

// Register for register
func (h *HandlerImpl) Register(c *gin.Context) {
	req, _ := c.Get("body")
	registerReq := req.(*dto.RegisterRequest)

	token, err := h.service.Register(registerReq.Username, registerReq.Name, registerReq.Password, registerReq.ConfirmPassword)
	if err != nil {
		h.responseHelper.ErrorResponse(c, http.StatusBadRequest, err.Error(), err.Error())
		return
	}

	h.responseHelper.SuccessResponse(c, http.StatusOK, "Success regist", token)

}
