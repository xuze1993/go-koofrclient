package auth

import (
	"fmt"

	"github.com/koofr/go-httpclient"
)

var (
	NotInitializedErr = fmt.Errorf("AuthProvider is not initialized")
)

type AuthProvider interface {
	Authenticate() error
	SetClient(c *httpclient.HTTPClient)
}