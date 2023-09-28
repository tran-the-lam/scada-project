package middleware

import (
	"fmt"
	"strings"

	"backend/pkg/constant"
	e "backend/pkg/error"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func userAuth(c *fiber.Ctx) error {
	reqToken := c.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	if len(reqToken) == 0 {
		return e.Unauthorized()
	}

	token, _ := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(constant.TOKEN_SECRET), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Locals(constant.LOCAL_USER_ID, claims["user_id"])
		c.Locals(constant.LOCAL_USER_ROLE, claims["user_role"])
		c.Locals(constant.LOCAL_IP_ADDR, c.Get("ip-addr"))
		c.Locals(constant.LOCAL_DEVICE_ID, c.Get("device-id"))
		c.Locals(constant.LOCAL_USER_AGENT, c.Get("user-agent"))
		return c.Next()
	} else {
		return e.Unauthorized()
	}
}

func sensorAuth(c *fiber.Ctx) error {
	// Todo Validate api key in here
	return nil
}

func Auth(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	if len(headers["Authorization"]) > 0 {
		return userAuth(c)
	} else if len(headers["X-Scada-Api-Key"]) > 0 {
		return sensorAuth(c)
	}

	return e.Unauthorized()
}
