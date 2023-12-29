package models

import "time"

const (
	expireationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expireationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expireationTime := time.Unix(at.Expires, 0)

	return expireationTime.Before(now)
}
