package db

import (
	"fmt"
	"log"

	"github.com/QR_code/models"
	"github.com/labstack/echo/v4"
)

func (db *QRAppDBImpl) UploadFormDB(c echo.Context, formData *models.UPlodFormResponse) error {
	tx := db.conn.MustBegin()
	_, err := tx.NamedQuery(`INSERT INTO "qr".form_data(name,image,qr_code)VALUES(:name,:image,:qr_code)`, formData)
	if err != nil {
		fmt.Print("cannot inserted row", err)
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Println("here is the error:", err)
		log.Fatal(err)
	}
	return nil
}
