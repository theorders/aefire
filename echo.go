package aefire

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"strings"
)

func NewAEFireEcho(aef *AEFire) *echo.Echo {
	e := NewEcho()
	e.Use(aef.EchoContextMiddleware)

	return e
}

func NewEcho() *echo.Echo {
	eco := echo.New()

	loggerConfig := middleware.DefaultLoggerConfig
	loggerConfig.Format = `{
   "time":"${time_rfc3339}",
   "method":"${method}",
   "uri":"${uri}",
   "status":"${status}",
   "error":"${error}",
}`

	eco.HTTPErrorHandler = func(err error, c echo.Context) {
		a := CastEchoContext(c)

		he, ok := err.(*echo.HTTPError)
		if ok {
			if he.Internal != nil {
				if herr, ok := he.Internal.(*echo.HTTPError); ok {
					he = herr
				}
			}
		} else {
			he = &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
		if eco.Debug {
			he.Message = err.Error()
		} else if m, ok := he.Message.(string); ok {
			he.Message = echo.Map{"message": m}
		}

		// Send response
		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead { // Issue #608
				err = c.NoContent(he.Code)
			} else {
				body := MapOf("message", he.Message)

				if a.UID() != nil {
					body["uid"] = *a.UID()
				} else {
					body["uid"] = "anonymous"
				}

				err = c.JSON(he.Code, body)
			}
			if err != nil {
				eco.Logger.Error(err)
			}
		}
	}

	eco.Use(
		middleware.Recover(), // Recover from all panics to always have your server up
		//middleware.LoggerWithConfig(loggerConfig),
		middleware.RequestID(), // Generate a request id on the HTTP response headers for identification
	)

	eco.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		a := CastEchoContext(c)

		req := MapOf()
		res := MapOf()

		if len(reqBody) > 0 {
			json.Unmarshal(reqBody, &req)
		}

		if len(resBody) > 0 && strings.Contains(c.Response().Header().Get("content-type"), "json") {
			json.Unmarshal(resBody, &res)
		}

		log := []interface{}{}

		if a.UID() != nil {
			log = append(log, *a.UID())
		} else {
			log = append(log, "anonymous")
		}

		log = append(log, c.Response().Status,
			c.Request().Method,
			c.Request().URL.String(),
			req,
			res,
			c.Request().Header.Get("User-Agent"))

		println(ToJson(log))
	}))

	if l, ok := eco.Logger.(*log.Logger); ok {
		//"${time_rfc3339_nano}","level":"${level}","prefix":"${prefix}","file":"${short_file}","line":"${line}"
		l.SetHeader(" ")
	}

	eco.Debug = true
	return eco
}

func JsonMap(c echo.Context, pairs ...interface{}) error {
	return c.JSON(200, MapOf(pairs...))
}

func JsonOk(c echo.Context) error {
	return c.JSON(200, MapOf("message", "ok"))
}

func JsonMsg(c echo.Context, msg string) error {
	return c.JSON(200, MapOf("message", msg))
}

func (aef *AEFire) IdTokenAuth(key string, c echo.Context) (result bool, err error) {
	a := aef.WithEcho(c)

	a.UserToken, err = a.Auth.VerifyIDToken(c.Request().Context(), key)

	if err == nil {
		return true, nil
	} else {
		return false, echo.NewHTTPError(401, "인증이 실패했습니다.")
	}
}

func (aef *AEFire) ValidateIdToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		a := aef.WithEcho(c)
		idToken := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")

		if idToken != "" {
			a.UserToken, _ = a.Auth.VerifyIDToken(c.Request().Context(), idToken)
		}

		return next(c)
	}
}

func AuthCronMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !strings.Contains(c.Request().Header.Get("X-Forwarded-For"), "0.1.0.1") {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		return next(c)
	}
}
