package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"web-server/cmd/registration/registration"
	"web-server/pkg/config"
	"web-server/pkg/helper/tokenhelper"
	"web-server/pkg/models"
	"web-server/pkg/models/authmodels"
)

func SignUp(c echo.Context) error {

	adminToken, err := tokenhelper.GetTokenFromHeader(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "getTokenFromHeader failed"}})
	}
	if adminToken == "" {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "token empty"}})
	}

	configs := config.GetConfig()
	if adminToken != configs.ConfigToken {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "admin tokens don't match"}})
	}

	user := &authmodels.User{}
	err = c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "bind failed"}})
	}

	if user.Login == "" || user.Password == "" {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "sorry, no credentials provided"}})
	}

	_, err = registration.SignUp(user.Login, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: fmt.Sprintf("signup failed: %v", err)}})
	}

	return c.JSON(http.StatusOK, models.Model{Response: &models.Response{Login: user.Login}})
}
