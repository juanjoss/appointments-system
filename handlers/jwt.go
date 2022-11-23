package handlers

import (
	"appointments-system/models"
	"appointments-system/ports"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CreateJWT(c *fiber.Ctx, admin models.Administrator) (string, time.Time, error) {
	claims := jwt.RegisteredClaims{
		ID:        admin.Email,
		Issuer:    admin.FullName,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", claims.ExpiresAt.Time, c.RedirectToRoute("login", ports.ErrorParam(ports.SignedStringError))
	}

	return tokenStr, claims.ExpiresAt.Time, nil
}

func GetJWT(c *fiber.Ctx) (*jwt.Token, *jwt.RegisteredClaims, error) {
	token := &jwt.Token{}
	claims := &jwt.RegisteredClaims{}

	tokenString := c.Cookies("jwt", "none")
	if tokenString == "none" {
		return token, claims, c.RedirectToRoute("login", nil)
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (any, error) { return []byte(GetSecret()), nil },
	)
	if err != nil || !token.Valid {
		return token, claims, Logout(c)
	}

	return token, claims, nil
}
