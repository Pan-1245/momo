package config

type Config struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

func NewConfig() *Config {
	return &Config{
		Prefix: "!momo",
	}
}
