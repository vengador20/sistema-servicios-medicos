package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vengador20/sistema-servicios-medicos/config"
)

func ValidateJwt(c *fiber.Ctx) error {

	//cook := c.Cookies("Authorized")

	// hed := c.GetReqHeaders()

	// fmt.Printf("hed: %v\n", hed)

	// ck := c.Request().Header.Cookie("token")

	// fmt.Printf("ck: %v\n", ck)

	cookie := c.Cookies("token") //Request().Header.Cookie("Authorized")

	fmt.Printf("cookie: %v\n", cookie)

	//fmt.Printf("cook: %v\n", cook)

	status, err := config.VerifyToken(string(cookie))

	if err != nil {
		return c.SendStatus(400)
	}

	if !status {
		return c.SendStatus(400)
	}

	return c.Next()
}
