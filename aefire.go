package aefire

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/iid"
	"firebase.google.com/go/messaging"
	"firebase.google.com/go/storage"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"google.golang.org/api/option"
)

type AEFire struct {
	context.Context
	App     *firebase.App
	FStore  *firestore.Client
	Storage *storage.Client
	FCM     *messaging.Client
	Auth    *auth.Client
	IID     *iid.Client

	DB        *sqlx.DB
	dbParam   DatabaseParam
	UserToken *auth.Token
}

func NewWithConfig(config *firebase.Config, opts ...option.ClientOption) *AEFire {
	return WithContext(context.Background(), config, opts...)
}

func New(opts ...option.ClientOption) *AEFire {
	return WithContext(context.Background(), nil, opts...)
}

func (aef *AEFire) ConnectDB(param DatabaseParam) *AEFire {
	aef.dbParam = param

	if aef.DB != nil {
		aef.DB.Close()
		aef.DB = nil
	}

	aef.DB = sqlx.MustConnect(param.DriverName, param.Url)

	PanicIfError(aef.DB.Ping())

	aef.DB.SetMaxOpenConns(param.MaxOpenConns)
	aef.DB.SetMaxIdleConns(param.MaxIdleConns)

	return aef
}

func WithContext(c context.Context, config *firebase.Config, opts ...option.ClientOption) *AEFire {
	aef := AEFire{}
	aef.Context = c

	var err error

	aef.App, err = firebase.NewApp(c, config, opts...)
	PanicIfError(err)

	aef.FStore, err = aef.App.Firestore(c)
	PanicIfError(err)

	aef.Storage, err = aef.App.Storage(c)
	PanicIfError(err)

	aef.FCM, err = aef.App.Messaging(c)
	PanicIfError(err)

	aef.Auth, err = aef.App.Auth(c)
	PanicIfError(err)

	aef.IID, err = aef.App.InstanceID(c)
	PanicIfError(err)

	return &aef
}

func (aef *AEFire) WithEcho(e echo.Context) *AEFire {
	return aef.WithContext(e.Request().Context())
}
func (aef *AEFire) WithContext(c context.Context) *AEFire {
	if a := CastContext(c); a != nil {
		return a
	}

	withCont := AEFire{}
	withCont = *aef
	withCont.Context = c

	if withCont.DB != nil && withCont.DB.Ping() != nil {
		withCont.ConnectDB(withCont.dbParam)
		aef.DB = withCont.DB
	}

	return &withCont
}

func (aef *AEFire) EchoContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		a := aef.WithEcho(e)

		e.SetRequest(e.Request().WithContext(a))

		return next(e)
	}
}

func CastContext(c context.Context) *AEFire {
	if a, ok := c.(*AEFire); ok {
		return a
	} else {
		return nil
	}
}

func CastEchoContext(e echo.Context) *AEFire {
	if a, ok := e.Request().Context().(*AEFire); ok {
		return a
	} else {
		return nil
	}
}
