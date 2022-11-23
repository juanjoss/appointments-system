package handlers

import (
	"appointments-system/db"
	"appointments-system/ports"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var secret = []byte("t1m0-ap7o1n7m3n7s")
var cookieKey = "jwt"

/*
Get the secret used for signing tokens.
*/
func GetSecret() []byte {
	return secret
}

/*
Authentication middleware for protecting administrator routes.
*/
func Auth(c *fiber.Ctx) error {
	_, _, err := GetJWT(c)
	if err != nil {
		return err
	}

	return c.Next()
}

/*
Handler to prevent accesing the login page after authentication.
*/
func LoginPage(c *fiber.Ctx) error {
	if c.Cookies(cookieKey, "none") != "none" {
		return c.RedirectToRoute("admin_appointments", nil)
	}

	return c.Render("auth/login", fiber.Map{
		"error": c.Query("error"),
	})
}

/*
Login logic to authenticate the administrator.
*/
func Login(c *fiber.Ctx) error {
	var request ports.LoginRequest

	// get request data
	err := c.BodyParser(&request)
	if err != nil {
		return c.RedirectToRoute(
			"login",
			ports.ErrorParam(ports.BodyParsingError),
		)
	}

	// get administrator
	admin, err := db.Get().GetAdmin()
	if err != nil {
		return c.RedirectToRoute(
			"login",
			ports.ErrorParam(ports.GetAdminError),
		)
	}

	// check email
	if admin.Email != request.Email {
		return c.RedirectToRoute(
			"login",
			ports.ErrorParam(ports.InvalidEmailError),
		)
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(request.Password))
	if err != nil {
		return c.RedirectToRoute(
			"login",
			ports.ErrorParam(ports.InvalidPasswordError),
		)
	}

	// create JWT
	token, exp, err := CreateJWT(c, admin)
	if err != nil {
		return err
	}

	// create and set cookie
	cookie := fiber.Cookie{
		Name:     cookieKey,
		Value:    token,
		Expires:  exp,
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.RedirectToRoute("admin_appointments", nil)
}

/*
Logout logic to close the administrator session.
*/
func Logout(c *fiber.Ctx) error {
	// not so elegant, but c.ClearCookie doesn't work
	c.Cookie(&fiber.Cookie{
		Name:     cookieKey,
		Value:    "",
		Expires:  time.Now(),
		HTTPOnly: true,
	})

	return c.RedirectToRoute("login", nil)
}
