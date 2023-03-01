package main

import (
	"bytes"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get email server configuration from environment variables
	server := os.Getenv("MAILTRAP_SERVER")
	port := os.Getenv("MAILTRAP_PORT")
	email := os.Getenv("MAILTRAP_EMAIL")
	password := os.Getenv("MAILTRAP_PASSWORD")

	// Set up Gin router
	router := gin.Default()

	// Define middleware to send email with log details
	router.Use(func(c *gin.Context) {
		// Log request using gin.Logger middleware
		gin.Logger()(c)

		// Compose email message with log details
		logMessage := bytes.Buffer{}
		logMessage.WriteString("Request URL: " + c.Request.URL.String() + "\n")
		logMessage.WriteString("Request Method: " + c.Request.Method + "\n")
		logMessage.WriteString("Request User-Agent: " + c.Request.UserAgent() + "\n")
		logMessage.WriteString("Request IP: " + c.ClientIP() + "\n")
		logMessage.WriteString("Error Log: " + c.Errors.String())

		auth := smtp.PlainAuth("", email, password, server)

		// Send email message
		to := []string{"kwamekyeimonies@gmail.com"} // Replace with actual email address
		msg := []byte("To: " + to[0] + "\r\n" +
			"Subject: Gin Log\r\n" +
			"\r\n" +
			logMessage.String())

		err := smtp.SendMail(server+":"+port, auth, email, to, msg)
		if err != nil {
			panic(err)
		}
	})

	// Define route handlers
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Start the server
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
