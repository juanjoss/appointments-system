package utils

import (
	"os"

	"gopkg.in/gomail.v2"
)

const (
	host = "smtp.gmail.com"
	port = 587

	bookingSuccessMsg = `
		<h1 align="center">¡Gracias!</h1>
		<br>

		<p>Tu reserva fue registrada para el día <i>%s %s de %d:00 a %d:00 horas</i>.</p>
		<br>

		<ul>
			<li>Tipo de Consulta: %s</li>
			<li>Método de Pago: %s</li>
			<li>Monto: $%.2f</li>
		</ul>

		<b><i>
		Para hacer una reprogramación o cancelacion de tu turno deberás ponerte en contacto, 
		al menos 24 horas antes de la fecha establecida, al teléfono %s o al email %s.
		</i></b>
		<br>

		<p>Esto es un email automatizado y no debes responder al mismo.</p>
	`

	bookingCancelledMsg = `
		<h1 align="center">¡Hola!</h1>
		<br>

		<p>
			Tu reserva para el día <i>%s %s de %d:00 a %d:00 horas</i> fue cancelada exitosamente. 
			Ante cualquier duda podés contactarte al teléfono %s o por email a %s.
		</p>
		<br>

		<p>Esto es un email automatizado y no debes responder al mismo.</p>
	`
)

var (
	username = os.Getenv("SENDER_EMAIL")
	password = os.Getenv("SENDER_PASSWORD")
)

func BookingSuccessTemplate() string {
	return bookingSuccessMsg
}

func BookingCancelledTemplate() string {
	return bookingCancelledMsg
}

func SendEmail(from, to, subject, body string) error {
	// build email
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(host, port, username, password)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
