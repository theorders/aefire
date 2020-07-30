package aefire

import (
	"net/url"
	"strconv"
)

func UrlValuesOf(args ...string) url.Values {
	param := url.Values{}

	if len(args)%2 != 0 {
		panic("Odd parameter count : " + strconv.Itoa(len(args)))
	}

	for i := 0; i < len(args); i += 2 {
		param.Set(args[i], args[i+1])
	}

	return param
}
