package utils

import (
	"crypto/tls"
	"github.com/Webservice-Rathje/Mailing-Service/config"
	"gopkg.in/gomail.v2"
)

func InitDialer() gomail.SendCloser {
	cfg, _ := config.ParseConfig()
	d := gomail.NewDialer(cfg.SMTP_Host, cfg.SMTP_Port, cfg.SMTP_Username, cfg.SMTP_Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}
	return s
}
