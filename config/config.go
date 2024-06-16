package config

type Config struct {
	Prefix string `json:"prefix"`
}

func NewConfig() *Config {
	return &Config{
		Prefix: "!momo",
	}
}
