package helpers

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
)

/**
* function to send email with subject and content
 */
func SendMail(email string, subject string, content string) {
	go func() {
		m := gomail.NewMessage()
		m.SetHeader("From", os.Getenv("STMP_EMAIL_SENDER"))
		m.SetHeader("To", email)
		m.SetHeader("Subject", os.Getenv("APP_NAME")+" "+subject)
		m.SetBody("text/html", content)

		// Send the email to Bob
		d := gomail.NewPlainDialer(os.Getenv("STMP_EMAIL_HOST"), 587, os.Getenv("STMP_EMAIL_ADDRESS"), os.Getenv("STMP_EMAIL_PASSWORD"))
		fmt.Println("++++++++++++")
		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
		fmt.Println("++++++++++++")
	}()
	return
}

//package helpers
//
//import (
//	"net/smtp"
//	"os"
//)
//
///**
//* function to send email with subject and content
// */
//func SendMail(email string, subject string, content string) {
//	go func() {
//		auth := smtp.PlainAuth("", os.Getenv("STMP_EMAIL_ADDRESS"), os.Getenv("STMP_EMAIL_PASSWORD"), os.Getenv("STMP_EMAIL_HOST"))
//		to := []string{email}
//		msg := []byte("To : " + email + "\r\n" + "Subject : " + os.Getenv("APP_NAME") + " " + subject + "\r\n" + "\r\n" + content + "\r\n" )
//		err := smtp.SendMail(os.Getenv("STMP_EMAIL_HOST")+":"+os.Getenv("STMP_EMAIL_PORT"), auth, os.Getenv("STMP_EMAIL_ADDRESS"), to, msg)
//		if err != nil {
//			return
//		}
//	}()
//	return
//}
