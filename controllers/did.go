package controllers

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/didbot/didbot-api-go/models"
)

func GetDids(c echo.Context) error {

	cursor  := &models.Cursor{}
	d 	:= &models.Did{}

	userId 	:= c.Get("token").(*models.Token).UserID
	q 	:= c.QueryParam("q")

	collection := d.GetDids(userId, q)
	cursor.SetCursor(nil, nil, collection.LastId(), collection.Len(), 20)
	r := FormatResponse(collection.Dids, cursor)

	return c.JSON(http.StatusOK, r)
}

func ShowDid(c echo.Context) error {
	_ = c.Param("id")
	return c.NoContent(http.StatusOK)
}

func CreateDid(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func DeleteDid(c echo.Context) error {
	_ = c.Param("id")
	return c.NoContent(http.StatusOK)
}