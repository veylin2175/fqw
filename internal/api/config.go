package api

type Config struct {
	Port   string
	APIKey string
}

func DefaultConfig() *Config {
	return &Config{
		Port:   "8080",
		APIKey: "your-secret-api-key-change-me",
	}
}
