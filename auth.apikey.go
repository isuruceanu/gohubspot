package gohubspot

import (
	"fmt"
	"net/http"
	"net/url"
)

type APIKeyAuth struct {
	apiKey string
}

const (
	apiKeyParam = "/?hapikey=%s"
)

// NewAPIKeyAuth create new API KEY Authenticator
func NewAPIKeyAuth(apikey string) APIKeyAuth {
	return APIKeyAuth{apiKey: apikey}
}

// Authenticate set auth
func (auth APIKeyAuth) Authenticate(request *http.Request) error {

	urlStr := request.URL.String() + fmt.Sprintf(apiKeyParam, auth.apiKey)
	url, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	request.URL = url

	return nil
}
