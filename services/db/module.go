package db

import (
	"fmt"

	"github.com/QR_code/models"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mubashir"
	password = "secret"
	dbname   = "qr_db"
	sslmode  = "disable"
)

type QRAppDB interface {
	UploadFormDB(c echo.Context, formData *models.UPlodFormResponse) error
	GetAllFormDB(c echo.Context) ([]*models.GetAllFormResponse, error)
}
type QRAppDBImpl struct {
	conn *sqlx.DB
}

func NewQRAppDBImpl() *QRAppDBImpl {
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", user, password, host, port, dbname, sslmode)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return &QRAppDBImpl{
		conn: db,
	}

}

var _ QRAppDB = &QRAppDBImpl{}
