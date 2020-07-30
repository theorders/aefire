package aefire

import (
	"github.com/labstack/echo/v4"
	"log"
)

func NewHttpError(params ...interface{}) *echo.HTTPError {
	he := echo.HTTPError{
		Code: 500,
	}

	for _, v := range params {
		switch vv := v.(type) {
		case int:
			he.Code = vv
		case string:
			he.Message = vv
		case echo.HTTPError:
			he = vv
		case *echo.HTTPError:
			he = *vv
		case error:
			if he.Message == nil {
				he.Message = vv.Error()
			}
			he.Internal = vv
		}
	}

	return &he
}

func PanicIfError(vs ...interface{}) {
	for i := len(vs) - 1; i > -1; i-- {
		v := vs[i]
		if v == nil {
			continue
		}

		switch vv := v.(type) {
		case error:
			println(vv.Error())
			panic(vv)
		}
	}
}

func LogIfError(vs ...interface{}) bool {
	for i := len(vs) - 1; i > -1; i-- {
		v := vs[i]
		if v == nil {
			continue
		}

		switch vv := v.(type) {
		case error:
			log.Println(vv.Error())
			return true
		}
	}

	return false
}
