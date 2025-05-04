package utils

import (
	"fmt"
	"net/smtp"
)

func SendEmailOTP(recipient, otp string) error {
	from := "andrewwilliamstaines@gmail.com"
	password := "ciwp dbks obmo agog"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	subject := "Subject: Mindly Register New\n"
	body := fmt.Sprintf("Your OTP for verification is: %s", otp)
	message := []byte(subject + "\n" + body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{recipient}, message)
	return err
}

func SendForgotMailOTP(recipient, otp string) error {
	from := "andrewwilliamstaines@gmail.com"
	password := "ciwp dbks obmo agog"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	subject := "Subject: Mindly Register New\n"
	body := fmt.Sprintf("Your OTP for verification is: %s", otp)
	message := []byte(subject + "\n" + body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{recipient}, message)
	return err
}
