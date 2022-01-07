package products

import (
	"fmt"
	response "github.com/WahidinAji/web-response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (d *Dependency) GetAll(ctx echo.Context) error {
	rows, err := d.FindAll(ctx.Request().Context())
	if err != nil {
		return err
	}
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK", rows))
}

func (d *Dependency) GetById(ctx echo.Context) error {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err.Error())
	}

	row, err := d.FindId(ctx.Request().Context(), postId)
	fmt.Println("row : ", row, " err : ", err)

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.WebResponse(http.StatusNotFound, "Not Found", err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK", row))
}

func (d *Dependency) UpdateById(ctx echo.Context) error {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product := new(Product)
	if err = ctx.Bind(product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(product); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.WebResponse(http.StatusBadRequest, "Bad Request", err.Error()))
	}

	row, err := d.Update(ctx.Request().Context(), postId, *product)

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.WebResponse(http.StatusNotFound, "Not Found", err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK", row))
}

func (d *Dependency) DeleteById(ctx echo.Context) error {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}
	err = d.Delete(ctx.Request().Context(), postId)

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.WebResponse(http.StatusNotFound, "Not Found", err.Error()))
	}
	return ctx.JSON(http.StatusOK, MsgDel{Code: 200, Status: "OK"})
}

func (d *Dependency) CreateOne(ctx echo.Context) error {
	product := new(Product)
	if err := ctx.Bind(product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(product); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.WebResponse(http.StatusBadRequest, "Bad Request", err.Error()))
	}

	row, err := d.Save(ctx.Request().Context(), *product)

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.WebResponse(http.StatusInternalServerError, "Internal Server Error", err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK", row))
}
