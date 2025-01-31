package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	secretKey := []byte("secret")

	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "missing token"})
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "invalid token format"})
		}

		tokenString := tokenParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, c.JSON(http.StatusInternalServerError, map[string]string{"message": "invalid signing method"})
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusInternalServerError, map[string]string{"messagae": "invalid claims"})
		}

		c.Set("user_email", claims["email"])
		c.Set("is_admin", claims["admin"])

		return next(c)
	}
}
