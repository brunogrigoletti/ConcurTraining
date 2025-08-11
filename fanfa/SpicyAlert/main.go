package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"net/url"
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
		sendPushNotification()
		sendTelegramNotification()
	} else {
		fmt.Println("No job yet... Keep looking!")
	}
}

func sendEmail() {
	//Sender Data
	from := "fanfathi95@gmail.com"
	password := os.Getenv("EMAIL_PASSWORD")

	//Receiver email address
	to := "fanfathi95@gmail.com"

	//SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	//Message
	message := []byte("Subject: !!!SpicyAlert!!! - Customer Love Engineer is Open!\r\n\r\nThe 'Customer Love Engineer' role is now open at Chili Piper. Go apply now!\r\nhttps://www.chilipiper.com/careers#work\r\n\r\nNão esquece de apagar a App-Password do código pra ninguém te hackear!!!")

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

func sendPushNotification() {
	topic := "spicyalert-thiago"
	message := "The 'Customer Love Engineer' role is now open at Chili Piper. Go apply now!\r\nhttps://www.chilipiper.com/careers#work"

	url := fmt.Sprintf("https://ntfy.sh/%s", topic)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(message)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Title", "!!!SpicyAlert!!! - Customer Love Engineer is Open!")
	req.Header.Set("Priority", "high")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending push notification:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted {
		fmt.Println("Push notification sent successfully!")
	} else {
		fmt.Println("Error sending push notification:", resp.Status)
	}
}

func sendTelegramNotification() {
	botToken := os.Getenv("TELEGRAM_TOKEN")
	chatID := "1001092787"
	message := "The 'Customer Love Engineer' role is now open at Chili Piper. Go apply now!\r\nhttps://www.chilipiper.com/careers#work"

	telegramAPI := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	resp, err := http.PostForm(telegramAPI, url.Values{
		"chat_id": {chatID},
		"text":    {message},
	})
	if err != nil {
		fmt.Println("Error sending Telegram notification:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Telegram notification sent successfully!")
	} else {
		fmt.Println("Error sending Telegram notification:", resp.Status)
	}
}
