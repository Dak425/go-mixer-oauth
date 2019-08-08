package mixeroauth

import (
	"time"
)

type Token struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	ExpiresOn    time.Time `json:"expires_on,omitempty"`
}
