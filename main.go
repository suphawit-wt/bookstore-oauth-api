package main

import (
	"bookstore/oauth/handler"
	"bookstore/oauth/repository"
	"bookstore/oauth/services"

	"github.com/gin-gonic/gin"
)

func main() {

	accessTokenRepositoryDB := repository.NewAccessTokenRepositoryDB()
	accessTokenService := services.NewAccessTokenService(accessTokenRepositoryDB)
	accessTokenHandler := handler.NewAccessTokenHandler(accessTokenService)

	router := gin.Default()

	router.GET("/oauth/access_token/:access_token_id", accessTokenHandler.GetAccessToken)

	router.Run(":8080")
}
