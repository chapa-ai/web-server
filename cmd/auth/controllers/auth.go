package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"web-server/pkg/db/authdb"
	"web-server/pkg/models"
	"web-server/pkg/models/authmodels"
	"web-server/pkg/models/documentsModels"
)

func Login(c echo.Context) error {
	user := &authmodels.User{}
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "bind failed"}})
	}

	userDb, err := authdb.GetUserAndPassword(c.Request().Context(), user.Login)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "getUserAndPassword failed"}})
	}
	if userDb == nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "error with credentials"}})
	}

	if !user.ComparePassword(userDb.Password) {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 401, Text: "compare password failed:"}})
	}

	tokenStr := user.CreateToken()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "createToken failed"}})
	}

	session := &authmodels.Session{
		UserId: userDb.Id,
		Token:  tokenStr,
	}
	_, err = authdb.SessionCreate(c.Request().Context(), session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: fmt.Sprintf("sessionCreate failed: %v", err)}})
	}

	return c.JSON(http.StatusOK, &documentsModels.Document{Token: tokenStr})
}

func Logout(c echo.Context) error {
	token := c.Param("token")
	user := &authmodels.User{Token: token}

	deletedToken, err := authdb.SessionDeleteToken(c.Request().Context(), user.Token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: fmt.Sprintf("sessionDeleteToken failed: %v", err)}})
	}

	return c.JSON(http.StatusOK, &models.Response{Response: map[string]bool{deletedToken: true}})
}
