package server

import (
	"github.com/QR_code/services/api"
	"github.com/labstack/echo/v4"
)

type ServerImpl interface {
	UploadForm(c echo.Context) error
	GetAllForms(c echo.Context) error
}
type Server struct {
	api *api.QRAppApiImpl
}

func NewServer() *Server {
	return &Server{
		api: api.NewQRAppApiImpl(),
	}
}
func NewServerImpl(e *echo.Echo) {
	server := NewServer()
	e.POST("/uploadform", server.UploadForm)
	e.GET("/getallforms", server.GetAllForms)
}

var _ ServerImpl = &Server{}
