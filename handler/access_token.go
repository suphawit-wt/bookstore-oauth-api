package handler

import (
	"bookstore/oauth/services"

	"github.com/gin-gonic/gin"
)

type accessTokenHandler struct {
	atSrv services.AccessTokenService
}

func NewAccessTokenHandler(atSrv services.AccessTokenService) accessTokenHandler {
	return accessTokenHandler{atSrv: atSrv}
}

func (h accessTokenHandler) GetAccessToken(c *gin.Context) {

}
