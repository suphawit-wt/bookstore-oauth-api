package repository

type AccessTokenRepository interface {
}

type accessTokenRepositoryDB struct{}

func NewAccessTokenRepositoryDB() AccessTokenRepository {
	return accessTokenRepositoryDB{}
}
