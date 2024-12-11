package auth

import (
	"os"
	"strings"
	"time"
	"golang.org/x/crypto/bcrypt"
	"cyberrange/db"
	"cyberrange/types"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
    var user types.UserLogin
    if err := c.Bind(&user); err != nil {
        return c.JSON(400, map[string]string{"error": "Invalid input"})
    }

    if user.EmailOrID == "" || user.Password == "" {
        return c.JSON(400, map[string]string{"error": "Please provide all required fields"})
    }

    user.EmailOrID = strings.ToLower(user.EmailOrID)

    var role, hashedPassword, name, email, user_id string
    query := "SELECT role, password, name, email, user_id FROM users WHERE (email = ? OR user_id = ?)"
    err := db.DB.QueryRow(query, user.EmailOrID, user.EmailOrID).Scan(&role, &hashedPassword, &name, &email, &user_id)
    if err != nil {
        return c.JSON(401, map[string]string{"error": "Failed to authenticate"})
    }

    if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)); err != nil {
        return c.JSON(401, map[string]string{"error": "Incorrect password"})
    }

    emailParts := strings.Split(email, "@")
    if len(emailParts) > 0 {
        email = emailParts[0]
    }

    token, err := generateToken(email, role, name, user_id)
    if err != nil {
        return c.JSON(500, map[string]string{"error": "Failed to generate token"})
    }

    return c.JSON(200, map[string]string{
        "token": token,
        "role":  role,
        "name":  name,
    })
}

func generateToken(id, role, name, user_id string) (string, error) {

	secretKey := []byte(os.Getenv("JWT_SECRET"))

	token, err := createToken(id, role, name, user_id, secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func createToken(id, role, name, user_id string, secretKey []byte) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["role"] = role
	claims["name"] = name
	claims["user_id"] = user_id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
