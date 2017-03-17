package controllers

import "github.com/didbot/didbot-api-go/models"

type Meta struct {
	Cursor  *models.Cursor  `json:"cursor"`
}

// Prepare response envelope
func FormatResponse(data interface{}, cursor  *models.Cursor) interface{}{
	r := struct {
		Data interface{}
		Meta interface{}
	}{
		data,
		Meta {
			cursor,
		},
	}

	return r
}
