package api

import (
	"fmt"
	"log"

	"github.com/QR_code/models"
	"github.com/skip2/go-qrcode"

	"github.com/labstack/echo/v4"
)

func (api *QRAppApiImpl) UploadFormAPI(c echo.Context, name string, image []byte) (*models.UPlodFormResponse, error) {
	var err error
	form_data := models.NewUPlodFormResponse()

	*form_data.Name = name
	*form_data.Image = image
	*form_data.QRCode, err = qrcode.Encode(*form_data.Name, qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
	}

	err = api.db.UploadFormDB(c, form_data)
	if err != nil {
		fmt.Println("API error while inserting data in db", err)
		log.Fatal(err)
	}

	return form_data, nil
}
