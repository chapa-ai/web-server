package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"web-server/pkg/db/documentsdb"
	"web-server/pkg/helper/filehelper"
	"web-server/pkg/helper/tokenhelper"
	"web-server/pkg/models"
	"web-server/pkg/models/documentsModels"
)

func CreateDocument(c echo.Context) error {
	meta := c.FormValue("meta")
	js := c.FormValue("json")

	data := &documentsModels.Document{}

	err := json.Unmarshal([]byte(meta), data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error: %v", err))
	}

	err = os.MkdirAll("temp", 0750)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "couldn't create directory"}})
	}

	f, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "formFile failed"}})
	}
	path := filepath.Join("temp", f.Filename)
	err = filehelper.CopyFileHeader2Str(f, path)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "copyFileHeader2Str failed"}})
	}

	err = documentsdb.InsertDocuments(data, js, path)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: fmt.Sprintf("insertDocuments failed: %v", err)}})
	}

	return c.JSON(http.StatusOK, models.Model{Data: documentsModels.Data{File: f.Filename, Json: data.Json}})
}

func GetDocument(c echo.Context) error {
	token, err := tokenhelper.GetTokenFromHeader(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "getTokenFromHeader failed"}})
	}
	if token == "" {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "token empty"}})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "strconv.Atoi failed"}})
	}

	documents := &documentsModels.Document{}
	err = c.Bind(documents)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "bind failed"}})
	}

	doc, err := documentsdb.GetDocument(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "getDocument failed"}})
	}
	if !documents.File {
		return c.JSON(http.StatusOK, models.Model{Data: doc})
	}

	path := filepath.Join(doc.Directory)
	return c.File(path)
}

func GetDocumentsList(c echo.Context) error {
	token, err := tokenhelper.GetTokenFromHeader(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "getTokenFromHeader failed"}})
	}
	if token == "" {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "token empty"}})
	}

	info := documentsModels.Pass{}
	err = c.Bind(&info)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "bind failed"}})
	}

	output, err := documentsdb.GetDocumentsList(c.Request().Context(), info)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: fmt.Sprintf("getDocumentsList failed: %v", err)}})
	}

	response := struct {
		Output []*documentsModels.Document `json:"docs"`
	}{Output: output}

	return c.JSON(http.StatusOK, models.Model{Data: response})
}

func DeleteDocument(c echo.Context) error {
	token, err := tokenhelper.GetTokenFromHeader(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "getTokenFromHeader failed"}})
	}
	if token == "" {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "token empty"}})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: "strconv.Atoi failed"}})
	}

	directory, err := documentsdb.DeleteDocument(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Model{Error: &models.Error{Code: 500, Text: fmt.Sprintf("deleteDocument failed: %v", err)}})
	}
	defer os.RemoveAll(directory)

	return c.JSON(http.StatusOK, &models.Response{Response: map[int]bool{id: true}})
}
