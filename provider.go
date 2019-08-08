package mixeroauth

import (
	"io"
	"net/http"
	"time"
)

const (
	clientIDHeader     = "Client-ID"
	clientSecretHeader = "Client-Secret"
	authHeader         = "Authorization"
	tokenPrefix        = "Bearer "
	grantTypeRefresh   = "refresh_token"
	grantTypeCode      = "authorization_code"
)

// Provider holds the refresh and access tokens used for Mixer's OAUTH2 gateway
type Provider struct {
	token  Token
	config Config
}

// NewProvider returns a pointer to a new provider of Mixer's OAUTH2 gateway
func NewProvider(config Config) Provider {
	return Provider{
		config: config,
	}
}

// Authenticated determines if the provider is currently authenticed with Mixer's OAUTH2 gateway
func (p *Provider) Authenticated() bool {
	return p.token.AccessToken != "" && p.token.ExpiresOn.Unix() > time.Now().Unix()
}

// func (p *Provider) Authenticate(client *http.Client) error {
// 	req := http.NewRequest(http.MethodPost, p.config.AuthURL)
// }

// func (p *Provider) Refresh() error {
// 	if p.token.RefreshToken == "" {
// 		return errors.New("no refresh token available")
// 	}

// }

// SetHeaders sets the authorization headers on the given http request
func (p *Provider) SetHeaders(req *http.Request) {
	if p.config.ClientID != "" {
		req.Header.Add(clientIDHeader, p.config.ClientID)
	}
	if p.config.ClientSecret != "" {
		req.Header.Add(clientSecretHeader, p.config.ClientSecret)
	}
	if p.Authenticated() {
		req.Header.Add(authHeader, tokenPrefix+p.token.AccessToken)
	}
}

// SetScopes indicates what permissions the auth token will grant the user
func (p *Provider) SetScopes(scopes ScopeCollection) {
	p.config.Scopes = scopes
}

func (p *Provider) request(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	p.SetHeaders(req)
	return req, nil
}

// func (p *Provider) refreshRequest() (*http.Request, error) {
// 	req, err := p.request(http.MethodPost)
// }
