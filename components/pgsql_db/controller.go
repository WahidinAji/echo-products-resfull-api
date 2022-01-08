package pgsql_db

import (
	"fmt"
	response "github.com/WahidinAji/web-response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (uv *UserValidator) Validate(i interface{}) error {
	if err := uv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (d *UserDependency) GetAll(ctx echo.Context) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rows, err := d.FindAll(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err.Error())
	}
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK", rows))
}

func (d *UserDependency) GetById(ctx echo.Context) error {
	userId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err.Error())
	}

	row, err := d.FindId(ctx.Request().Context(), userId)
	fmt.Println("row : ", row, " err : ", err)

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.WebResponse(http.StatusNotFound, "Not Found", err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK", row))
}

func (d *UserDependency) UpdateById(ctx echo.Context) error {
	userId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err.Error())
	}

	var user = new(User)
	if err = ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.WebResponse(http.StatusBadRequest, "Bad Request", err.Error()))
	}

	row, err := d.Update(ctx.Request().Context(), userId, *user)
	fmt.Println("row : ", row, " err : ", err)

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.WebResponse(http.StatusNotFound, "Not Found", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK", row))
}
