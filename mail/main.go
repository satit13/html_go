package main

import (
	"net/smtp"
	//"log"

	"fmt"
	"github.com/scorredoira/email"
	"log"
	_ "net/mail"
	//"github.com/matryer/m"
)

func main() {
	//m := email.NewMessage("Hi", "this is the body")
	m := email.NewHTMLMessage("Hi", "this is the body")
	m.From.Name = "go@nopadol.com"
	m.From.Address = "go mail"
	//m.To = []string{"tom@nopadol.com","satit@nopadol.com"}
	m.To = []string{"satit@nopadol.com"}
	m.Cc = []string{"it@nopadol.com"}
	m.Body = "test send main from golang"
	m.Subject = "send html mail from golang programming"

	err := m.Attach("./mail/gofer.png")
	err = m.Attach("1.html")
	// use Inline to display the attachment inline.
	if err := m.Inline("main.go"); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Println(err)
	}

	err = email.Send("smtp.gmail.com:587", smtp.PlainAuth("", "satit@nopadol.com", "815309917", "smtp.gmail.com"), m)
	if err != nil {
		fmt.Println("error send mail ", err.Error())
	}

}
