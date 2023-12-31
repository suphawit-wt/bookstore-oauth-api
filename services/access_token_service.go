package services

import (
	"bookstore/oauth/models"
	"bookstore/oauth/repository"
	"bookstore/oauth/utils/errors"
)

type AccessTokenService interface {
	GetById(string) (*models.AccessToken, *errors.RestErr)
}

type accessTokenService struct {
	atRepo repository.AccessTokenRepository
}

func NewAccessTokenService(atRepo repository.AccessTokenRepository) AccessTokenService {
	return accessTokenService{atRepo: atRepo}
}

func (s accessTokenService) GetById(token string) (*models.AccessToken, *errors.RestErr) {
	return nil, nil
}
