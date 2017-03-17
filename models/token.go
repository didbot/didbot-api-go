package models

import (
	//"gopkg.in/mgutz/dat.v2/dat"
	"github.com/satori/go.uuid"
)

type Token struct {
	Token     string    	`db:"id" json:"id"`
	UserID    uuid.UUID     `db:"user_id" json:"user_id"`
	ClientId  uuid.UUID     `db:"client_id" json:"client_id"`
	CreatedAt string        `db:"created_at" json:"created_at"`
}


func (t * Token) ValidateToken(tokenId string,userId string) (*Token, *error){
	token := Token{}
	err := DB.
	Select("user_id").
		From("oauth_access_tokens").
		Where("id = $1", tokenId).
		Where("user_id = $1", userId).
		Where("revoked = false").
		Where("expires_at >= NOW()").
		QueryStruct(&token)
	// TODO: Cache(tokenId, 30 * time.Second, false).
	// TODO: Validate client has not been revoked
	if err != nil {
		return nil, &err
	}

	return &token, nil
}