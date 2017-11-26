package main

import (
	"fmt"
	"net/smtp"
)

func enviarEmail(m string) {
	auth := smtp.PlainAuth("",
		"mail",
		"senha",
		"smtp.gmail.com",
	)

	smtp.SendMail("smtp.gmail.com:587",
		auth,
		"junioralberto19@gmail.com",
		[]string{"mail@gmail.com"},
		[]byte(m),
	)
	fmt.Printf("email %s enviado com sucesso", m)
}

func main() {
	fmt.Println("insira a mensagem do email")
	var m string
	fmt.Scanln(&m)

	go enviarEmail(m)

	fmt.Println("foi solicitado o envio do email, mas continuaremos com nosso programa ;)")

	fmt.Scanln(&m)
}
