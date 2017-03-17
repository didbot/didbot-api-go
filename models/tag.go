package models


import (
	//"gopkg.in/mgutz/dat.v2/dat"
	"github.com/satori/go.uuid"
)


type Tag struct {
	ID        uuid.UUID     `db:"id" json:"id"`
	UserID    uuid.UUID     `db:"user_id" json:"user_id"`
	Text      string      	`db:"text" json:"text"`
	CreatedAt string  	`db:"created_at" json:"created_at"`
}
