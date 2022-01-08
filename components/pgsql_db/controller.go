package pgsql_db

import (
	response "github.com/WahidinAji/web-response"
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
