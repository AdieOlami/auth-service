package controller

import (
	"net/http"

	"github.com/AdieOlami/auth-service/src/errors"
	"github.com/AdieOlami/auth-service/src/model"
	"github.com/AdieOlami/auth-service/src/services"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	AuthenticateRequest(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service services.Service
}

func NewController(service services.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) AuthenticateRequest(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request model.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	accessToken, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
