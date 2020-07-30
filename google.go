package aefire

import (
	"context"
	"github.com/imroc/req"
)

type GoogleAuthToken struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Azp string `json:"azp"`
	Aud string `json:"aud"`
	Iat string `json:"iat"`
	Exp string `json:"exp"`

	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
}

func VerifyGoogleToken(c context.Context, accessToken string) (info *GoogleAuthToken, err error) {
	res, err := req.Get(
		"https://www.googleapis.com/oauth2/v3/tokeninfo",
		UrlValuesOf("id_token", accessToken))

	if err != nil {
		return nil, err
	}

	info = &GoogleAuthToken{}
	if res.ToJSON(info) != nil {
		return nil, err
	}

	return info, nil
}
