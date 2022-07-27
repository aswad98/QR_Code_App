package api

import (
	"github.com/QR_code/models"
	"github.com/QR_code/services/db"
	"github.com/labstack/echo/v4"
)

type QRAppApi interface {
	UploadFormAPI(c echo.Context, name string, image []byte) (*models.UPlodFormResponse, error)
	GetAllFormsApi(c echo.Context) ([]*models.GetAllFormResponse, error)
}
type QRAppApiImpl struct {
	db *db.QRAppDBImpl
}

func NewQRAppApiImpl() *QRAppApiImpl {
	dbImpl := db.NewQRAppDBImpl()
	return &QRAppApiImpl{
		db: dbImpl,
	}
}

var _ QRAppApi = &QRAppApiImpl{}
