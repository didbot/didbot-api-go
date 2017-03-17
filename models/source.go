package models

import (
	"github.com/satori/go.uuid"
)


type Source struct {
	ID        uuid.UUID     `db:"id" json:"id"`
	Name      string      	`db:"name" json:"name"`
}