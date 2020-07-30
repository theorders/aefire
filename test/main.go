package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/theorders/aefire"
	"google.golang.org/api/option"
	"net/http"
)

var (
	key = []byte(``)
	aef = aefire.New(option.WithCredentialsJSON(key))
)

func main() {
	e := aefire.NewAEFireEcho(aef)

	cors := middleware.DefaultCORSConfig
	cors.AllowMethods = append(cors.AllowMethods, http.MethodOptions)

	e.GET("/", func(context echo.Context) error {
		panic(errors.New("oh my god"))
		return context.JSON(200, aefire.MapOf("message", "ok"))
	})

	aefire.PanicIfError(
		e.Start(":8080"))
}
