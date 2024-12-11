package auth

import (
	"strings"

	"cyberrange/db"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type ChangePassword struct {
	Email           string `json:"email"`
	Otp             string `json:"otp"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func ChangePass(c echo.Context) error {
	var changePassword ChangePassword

	err := c.Bind(&changePassword)
	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	if changePassword.Email == "" {
		return c.JSON(400, map[string]string{"error": "Please provide the email"})
	}

	if changePassword.Otp == "" {
		return c.JSON(400, map[string]string{"error": "Please provide the OTP"})
	}

	if changePassword.Password == "" {
		return c.JSON(400, map[string]string{"error": "Please provide the password"})
	}

	if changePassword.ConfirmPassword == "" {
		return c.JSON(400, map[string]string{"error": "Please provide the confirm password"})
	}

	if changePassword.Password != changePassword.ConfirmPassword {
		return c.JSON(400, map[string]string{"error": "Password and confirm password do not match"})
	}

	changePassword.Email = strings.ToLower(changePassword.Email)

	if len(changePassword.Password) < 8 {
		return c.JSON(400, map[string]string{"error": "Password must be at least 8 characters"})
	}

	var email string

	err = db.DB.QueryRow("SELECT email FROM users WHERE email = ?", changePassword.Email).Scan(&email)
	if err != nil {
		return c.JSON(401, map[string]string{"error": "User not found"})
	}

	var otp string

	err = db.DB.QueryRow("SELECT otp FROM users WHERE email = ?", changePassword.Email).Scan(&otp)
	if err != nil {

		return c.JSON(401, map[string]string{"error": "OTP is wrong"})

	}

	if otp != changePassword.Otp {
		return c.JSON(401, map[string]string{"error": "OTP is wrong"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changePassword.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to hash password"})
	}

	_, err = db.DB.Exec("UPDATE users SET password = ? WHERE email = ?", hashedPassword, changePassword.Email)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to update password"})
	}

	return c.JSON(200, map[string]string{"message": "Password updated successfully"})

}
