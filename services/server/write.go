package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/QR_code/models"
	"github.com/labstack/echo/v4"
)

func (s *Server) UploadForm(c echo.Context) error {

	name := c.FormValue("name")
	files, err := c.FormFile("image")
	if err != nil {
		log.Fatal(err)
	}
	fileDAta, err := files.Open()
	if err != nil {
		log.Fatal(err)
	}
	imageFile, err := ioutil.ReadAll(fileDAta)
	if err != nil {
		log.Fatal(err)
	}
	formData, err := s.api.UploadFormAPI(c, name, imageFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("data form", formData)

	result := models.Response{
		Data:    formData,
		Message: "Success"}

	return c.JSON(http.StatusOK, result)
}
