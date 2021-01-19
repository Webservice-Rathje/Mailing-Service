package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	SMTP_Host     string `json:"smtp_host"`
	SMTP_Port     int    `json:"smtp_port"`
	SMTP_Username string `json:"smtp_username"`
	SMTP_Password string `json:"smtp_password"`
}

func ParseConfig() (c *Config, err error) {
	f, err := os.Open("./config/config.json")
	if err != nil {
		return
	}
	c = new(Config)
	err = json.NewDecoder(f).Decode(c)
	return
}
