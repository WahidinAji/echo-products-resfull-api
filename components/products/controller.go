package products

import (
	response "github.com/WahidinAji/web-response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (d *Dependency) GetAll(ctx echo.Context) error {
	rows, err := d.FindAll(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK,"OK", rows))
}
