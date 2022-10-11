package tokenhelper

import (
	"fmt"
	"github.com/labstack/echo"
	"strings"
)

func GetTokenFromHeader(c echo.Context) (string, error) {
	auth := c.Request().Header["Authorization"]
	if len(auth) == 0 {
		return "", fmt.Errorf("Error getting token from header")
	}

	return strings.TrimPrefix(auth[0], "Bearer "), nil
}
