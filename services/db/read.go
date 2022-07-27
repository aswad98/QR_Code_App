package db

import (
	"log"

	"github.com/QR_code/models"
	"github.com/labstack/echo/v4"
)

func (db *QRAppDBImpl) GetAllFormDB(c echo.Context) ([]*models.GetAllFormResponse, error) {
	allForms := make([]*models.GetAllFormResponse, 0)
	err := db.conn.Select(&allForms, `select * from "qr".form_data;`)
	if err != nil {
		log.Fatal(err)
	}
	return allForms, err
}
