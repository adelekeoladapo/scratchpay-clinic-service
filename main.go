package main

import (
	"github.com/labstack/echo/v4"
	"fmt"
	"net/http"
)

func main() {
	e := echo.New()
	fmt.Println("Hello ScratchPay!")

	/* App Router */
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello ScratchPay")
	})
	e.Logger.Fatal(e.Start(":8090"))
}
