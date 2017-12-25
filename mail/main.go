package main

import (
	"net/smtp"
	//"log"

	"fmt"
	"github.com/scorredoira/email"
	"log"
	_ "net/mail"
)

func main() {
	m := email.NewMessage("Hi", "this is the body")
	m.From.Name = "go@nopadol.com"
	m.From.Address = "go mail"
	m.To = []string{"satit@extensionsoft.biz"}
	m.Body = "test send main from golang"
	m.Subject = "send main from golang"

	err := m.Attach("gofer.png")
	if err != nil {
		log.Println(err)
	}

	err = email.Send("smtp.gmail.com:587", smtp.PlainAuth("", "satit@nopadol.com", "815309917", "smtp.gmail.com"), m)
	if err != nil {
		fmt.Println("error send mail ", err.Error())
	}

}
