package mixeroauth

import "os"

const (
	oauth2AuthURL  = "https://mixer.com/oauth/authorize"
	oauth2TokenURL = "https://mixer.com/api/v1/oauth/token"
)

// Config contains the setting used to utilize Mixer's API
type Config struct {
	ClientID     string
	ClientSecret string
	AuthURL      string
	TokenURL     string
	Scopes       ScopeCollection
}

// NewConfig returns a new instance of a config used for the API
func NewConfig() Config {
	return Config{
		ClientID:     getEnv("MIXER_OAUTH2_CLIENT_ID", ""),
		ClientSecret: getEnv("MIXER_OAUTH2_CLIENT_SECRET", ""),
		AuthURL:      getEnv("MIXER_OAUTH2_AUTH_URL", oauth2AuthURL),
		TokenURL:     getEnv("MIXER_OAUTH2_TOKEN_URL", oauth2TokenURL),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
