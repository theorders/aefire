package aefire

import (
	"github.com/labstack/echo/v4"
	"testing"
)

func GetError() *echo.HTTPError {
	return nil
}

func TestNil(t *testing.T) {
	c := GetError()

	//if c == nil || (reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil())

	println(c == nil)
}
