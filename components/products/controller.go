package products

import (
	"fmt"
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
	fmt.Println("err : ",err, " row : ",row)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.WebResponse(http.StatusNotFound, "Not Found", err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK", row))
}

func (d *Dependency) UpdateById(ctx echo.Context) error {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,err.Error())
	}

	product := new(Product)
	if err = ctx.Bind(product); err != nil {
		return err
	}
	row, err := d.Update(ctx.Request().Context(), postId, *product)
	fmt.Println("err : ",err, " row : ",row)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.WebResponse(http.StatusNotFound, "Not Found s", err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK", row))
}


func (d *Dependency) DeleteById(ctx echo.Context) error {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}
	err = d.Delete(ctx.Request().Context(), postId)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.WebResponse(http.StatusNotFound, "Not Found", err.Error()))
	}
	return ctx.JSON(http.StatusOK, MsgDel{Code: 200, Status: "OK"})

}

func (d *Dependency) CreateOne(ctx echo.Context) error {
	product := new(Product)
	if err := ctx.Bind(product); err != nil {
		return err
	}
	row, err := d.Save(ctx.Request().Context(), *product)
	if err != nil {
		return err
	}
	if row == nil {
		return ctx.JSON(http.StatusNotFound, response.WebResponse(http.StatusNotFound, "Error", "Failed to created"))
	}
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK", row))
}

