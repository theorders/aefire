package aefire

/*
func (a *AEFire) MakeIdToken(apiKey, uid string) (verifAssertRes *identitytoolkit.VerifyAssertionResponse, err error) {
	ct, err := a.Auth.CustomToken(a, uid)

	if err != nil {
		return nil, err
	}

	idCli, err := identitytoolkit.NewService(a)
	if err != nil{
		return nil, err
	}

	idCli.Relyingparty.

	return identitytoolkit.SignInWithCustomToken(ods.Http(), apiKey, ct)
}

func (ods *ODSContext) IdTokenHeader(apiKey, uid string) (header req.Header, err error) {
	idToken, err := ods.MakeIdToken(apiKey, uid)

	if err != nil {
		return header, err
	}

	header = req.Header{"Authorization": "Bearer " + idToken.IdToken}

	return
}
*/
