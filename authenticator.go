package gohubspot

import (
	"net/http"
)

// Authenticator interface for auth
type Authenticator interface {
	Authenticate(request *http.Request) error
}
