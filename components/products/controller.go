package products

import (
	response "github.com/WahidinAji/web-response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (d *Dependency) GetAll(ctx echo.Context) error {
	rows, err := d.FindAll(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK", rows))
}
func (d *Dependency) GetById(ctx echo.Context) error {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}
	row, err := d.FindId(ctx.Request().Context(), postId)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK", row))
}

func (d *Dependency) UpdateById(ctx echo.Context) error {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}

	product := new(Product)
	if err = ctx.Bind(product); err != nil {
		return err
	}
	row, err := d.Update(ctx.Request().Context(), postId, *product)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK", row))
}
