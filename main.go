package main

func main() {
	from := "mymail@example.com"
	to := "to@mail.com"
	subject := "This is email subject"
	body_text := "Hello"
	body_html := "<b>Hello</b>"

	SendMailHTML(from, to, subject, body_text, body_html)
}
