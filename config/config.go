package config

import "os"

type Config struct {
	Port                   string
	DatabaseURL            string
	RedisURL               string
	CredentialEncryptionKey string
	ClaudeOAuthClientID    string
	ClaudeOAuthRedirectURI string
	ExportAPIKey           string
	ExportAPIAllowedIPs    string
}

func Load() *Config {
	return &Config{
		Port:                    envOr("PORT", "8080"),
		DatabaseURL:             os.Getenv("DATABASE_URL"),
		RedisURL:                os.Getenv("REDIS_URL"),
		CredentialEncryptionKey: os.Getenv("CREDENTIAL_ENCRYPTION_KEY"),
		ClaudeOAuthClientID:     os.Getenv("CLAUDE_OAUTH_CLIENT_ID"),
		ClaudeOAuthRedirectURI:  os.Getenv("CLAUDE_OAUTH_REDIRECT_URI"),
		ExportAPIKey:            os.Getenv("EXPORT_API_KEY"),
		ExportAPIAllowedIPs:     os.Getenv("EXPORT_API_ALLOWED_IPS"),
	}
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
