package gohubspot

import "net/http"

// OAuth2 OAuth 2 hubspot authenticator
type OAuth2 struct {
	token string
}

// NewOAuth2 Create new instance of OAuth2
func NewOAuth2(token string) OAuth2 {
	return OAuth2{token: token}
}

// Authenticate auth with OAuth2
func (auth OAuth2) Authenticate(request *http.Request) error {

	request.Header.Set("Authorization", "Bearer "+auth.token)

	return nil
}
