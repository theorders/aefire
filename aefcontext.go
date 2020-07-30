package aefire

import (
	"context"
	"firebase.google.com/go/auth"
)

func (a *AEFire) UID() *string {
	if a.UserToken == nil {
		return nil
	}

	return &a.UserToken.UID
}

func (a *AEFire) UserRecord() *auth.UserRecord {
	if a.UID() == nil {
		return nil
	}

	user, _ := a.Auth.GetUser(context.Background(), *a.UID())

	return user
}
