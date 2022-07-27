package api

import (
	"log"

	"github.com/QR_code/models"
	"github.com/labstack/echo/v4"
)

func (api *QRAppApiImpl) GetAllFormsApi(c echo.Context) ([]*models.GetAllFormResponse, error) {
	var getAllData []*models.GetAllFormResponse
	var err error
	getAllData, err = api.db.GetAllFormDB(c)
	if err != nil {
		log.Fatal(err)
	}
	return getAllData, nil
}
