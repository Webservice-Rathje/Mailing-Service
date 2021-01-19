package controller

import (
	"github.com/Webservice-Rathje/Mailing-Service/utils"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"log"
	"strings"
)

func SendRegistrationCodeController(c *fiber.Ctx) error {
	if c.IP() == "127.0.0.1" || c.IP() == "10.11.0.4" || c.IP() == "10.11.0.5" {
		mail := c.Query("mail")
		name := c.Query("name")
		two_fa_code := strings.Split(c.Query("2FA-Code"), "")
		s := utils.InitDialer()
		m := gomail.NewMessage()
		m.SetHeader("From", "noreply@webservice-rathje.de")
		m.SetAddressHeader("To", mail, name)
		m.SetHeader("Subject", "Webservice Rathje Kontoerstellung")
		data, _ := ioutil.ReadFile("./sample/RegistrationCode.html")
		content := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(string(data), "{{FIRST_CODE_PART}}", two_fa_code[0]+two_fa_code[1]+two_fa_code[2]), "{{FIRST_CODE_PART}}", two_fa_code[3]+two_fa_code[4]+two_fa_code[5]), "{{NAME}}", name)
		m.SetBody("text/html", content)
		if err := gomail.Send(s, m); err != nil {
			log.Printf("Could not send email to %q: %v", mail, err)
		}
		m.Reset()
		return c.SendString("Send Mail successfully")
	} else {
		return c.SendString("Your origin is not allowed")
	}
}
