package models

import (
	"github.com/satori/go.uuid"
)

type Collection interface {
	lastId() uuid.UUID
	len() int
}

type Cursor struct{
	Current   *string        `json:"current"`
	Prev      *string        `json:"prev"`
	Next      *string        `json:"next"`
        Count     int 		`json:"count"`
}

func (c *Cursor) SetCursor(current *string, prev *string, lastId *string, count int, limit int) {
	c.Current = current
	c.Prev    = prev
	c.Count   = count
	c.setNext(count, limit, lastId)
}

func (c *Cursor) setNext(count, limit int, lastId *string) {

	if count == 0 || limit > count {
		return
	}

	c.Next = lastId
}