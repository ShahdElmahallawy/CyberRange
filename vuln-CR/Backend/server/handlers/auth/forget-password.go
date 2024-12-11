package auth

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"cyberrange/db"

	"github.com/labstack/echo/v4"
)

type ForgetPassword struct {
	Email string `json:"email"`
}

func ForgetP(c echo.Context) error {
	var forgetPassword ForgetPassword

	err := c.Bind(&forgetPassword)
	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	if forgetPassword.Email == "" {
		return c.JSON(400, map[string]string{"error": "Please provide the email"})
	}

	forgetPassword.Email = strings.ToLower(forgetPassword.Email)

	var email string

	err = db.DB.QueryRow("SELECT email FROM users WHERE email = ?", forgetPassword.Email).Scan(&email)
	if err != nil {
		return c.JSON(401, map[string]string{"error": "User not found"})
	}

	otp := generateOTP()

	_, err = db.DB.Exec("UPDATE users SET otp = ? WHERE email = ?", otp, forgetPassword.Email)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to update OTP"})
	}

	return c.JSON(200, map[string]string{"message": "Check your email for OTP to reset your password"})

}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())

	otp := rand.Intn(900000) + 100000 // Generate a 6-digit OTP

	return fmt.Sprintf("%06d", otp)
}