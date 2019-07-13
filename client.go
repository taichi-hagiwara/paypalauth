package paypalauth

import (
	"fmt"

	"golang.org/x/oauth2"
)

// Client は、クライアント。
type Client struct {
	ClientID     string
	ClientSecret string
	Endpoint     Endpoint
}

// OAuth2Config は、 *oauth2.Config を得る。
func (c *Client) OAuth2Config(redirectURL string, scopes ...string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   c.Endpoint.AuthURL,
			TokenURL:  fmt.Sprintf("%s/v1/oauth2/token", c.Endpoint.RootURL),
			AuthStyle: oauth2.AuthStyleInHeader,
		},
		RedirectURL: redirectURL,
		Scopes:      scopes,
	}
}

// Token は、 *oauth2.Token から *Token を得る。
func (c *Client) Token(token *oauth2.Token) *Token {
	return &Token{
		Client:      c,
		OAuth2Token: token,
	}
}
