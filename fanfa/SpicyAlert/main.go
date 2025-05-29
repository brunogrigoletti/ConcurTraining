package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strings"
)

func main() {
	url := "https://www.chilipiper.com/careers#work"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	content := string(body)

	// Check if the job title exists in the page content
	if strings.Contains(content, "Customer Love Engineer") {
		fmt.Println("FOUND IT! GO APPLY!")
		//Send email alert
		sendEmail()
	} else {
		fmt.Println("No job yet... Keep looking!")
	}
}

func sendEmail() {
	//Sender Data
	from := "fanfathi95@gmail.com"
	password := "rvuu tkfk edyu oems"

	//Receiver email address
	to := "fanfathi95@gmail.com"

	//SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	//Message
	message := []byte("Subject: !!!SpicyAlert!!! - Customer Love Engineer is Open!\r\n\r\nThe 'Customer Love Engineer' role is now open at Chili Piper. Go apply now!\r\nhttps://www.chilipiper.com/careers#work \r\n\r\nNão esquece de apagar a App-Specific Password do código pra não te hackearem no futuro!!!")

	//Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	//Sending email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
	fmt.Println("Email sent successfully!")
}
