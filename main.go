package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"scratchpay.com/clinics/controller"
)


func main() {
	e := echo.New()
	fmt.Println("Hello ScratchPay!")

	/* Load configuration */
	err := godotenv.Load()
	if err != nil {
		log.Panic("Could not load configuration")
	}

	/* App Router */
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello ScratchPay")
	})
	e.GET("/api/v1/clinic/search", func(context echo.Context) error {
		return controller.Search(context)
	})
	e.Logger.Fatal(e.Start(":8090"))
}

