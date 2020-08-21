package tool

type MailHandler struct {
	user       string
	passwd     string
	smtpServer string
}

//func (m MailHandler) Send(subject string, body string, to []string) error {
//	auth := smtp.PlainAuth("", m.user, m.passwd, m.smtpServer)
//	headerSubject := fmt.Sprintf("Subject: %s\r\n", subject)
//	msg := []byte(headerSubject + "\r\n" + body)
//
//	err := smtp.SendMail()
//}
