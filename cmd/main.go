package main

import (
	"github.com/QR_code/services/server"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	server.NewServerImpl(e)
	e.Logger.Fatal(e.Start(":8080"))
}
