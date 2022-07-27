package server

import (
	"log"
	"net/http"

	"github.com/QR_code/models"
	"github.com/labstack/echo/v4"
)

func (s *Server) GetAllForms(c echo.Context) error {
	allForms, err := s.api.GetAllFormsApi(c)
	if err != nil {
		log.Fatal(err)
	}
	result := models.Response{
		Data:    allForms,
		Message: "success",
	}
	return c.JSON(http.StatusOK, result)

}
