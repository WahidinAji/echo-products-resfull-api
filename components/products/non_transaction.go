package products

import (
	"context"
	"fmt"
	response "github.com/WahidinAji/web-response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//Get All Data With No Transaction Process
func (d *Dependency) RepoProductsAll(ctx context.Context) ([]Product, error){
	db, err := d.DB.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.QueryContext(ctx,"SELECT id, name, stock, price FROM products")
	if err != nil {
		return nil, err
	}
	var products []Product
	for rows.Next(){
		var product Product
		err := rows.Scan(&product.ID,&product.Name,&product.Stock,&product.Price)
		if err != nil {
			return []Product{}, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (d *Dependency) ProductsAll(ctx echo.Context) error {
	data, err := d.RepoProductsAll(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK with no transaction proccess",data))
}

//Update Data By id With No Transaction Process

func (d *Dependency) RepoProductUpdate(ctx context.Context, id int64) (Product, error) {
	db, err := d.DB.Conn(ctx)
	if err != nil {
		return Product{}, err
	}
	defer db.Close()
	stmt, err := db.PrepareContext(ctx,"SELECT name, stock, price FROM products WHERE id=?")
	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()
	var product Product
	product.ID = int(id)
	product.Name = "Product One Update"
	product.Stock = 14
	product.Price = 15000.88
	sql := "update products SET name=?, stock=?,price=? where id=?"
	fmt.Println(id)
	res, err := d.DB.ExecContext(ctx, sql, &product.Name,&product.Stock,&product.Price, id)
	if err != nil {
		return Product{}, err
	}
	fmt.Println(res)
	return product, nil
}

func (d *Dependency) ProductUpdate(ctx echo.Context) error {
	productId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}
	data, err := d.RepoProductUpdate(ctx.Request().Context(), int64(productId))
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK with no transaction process updated success !!",data))
}

/**
Tebet jakarta Invesnit bisnis sinergi
-

Panorama Central Wisata, Jabar perbatasan Jakpus/Jakbar
-
 */