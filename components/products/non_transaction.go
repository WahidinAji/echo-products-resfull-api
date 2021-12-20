package products

import (
	"context"
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

func (d *Dependency) RepoProductUpdate(ctx context.Context, id int) (Product, error) {
	db, err := d.DB.Conn(ctx)
	if err != nil {
		return Product{}, err
	}
	defer db.Close()
	stmt, err := db.PrepareContext(ctx,"SELECT id, name, stock, price FROM products WHERE id=?")
	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()
	var name string
	stmt.QueryRowContext(ctx, id).Scan(&name)
	var product Product
	product.Name = name
	return product, nil

	//var product Product
	//rows, err := stmt.QueryContext(ctx,id)
	//if err != nil {
	//	return Product{}, nil
	//}
	//for rows.Next(){
	//	var p Product
	//	err = rows.Scan(&p.ID, &p.Name,&p.Stock,&p.Price)
	//	if err != nil {
	//		return Product{}, nil
	//	}
	//	product.ID = p.ID
	//	product.Name = "update"
	//	product.Stock = p.Stock
	//	product.Price = p.Price
	//	stmt.Query("SET name =?",product.Name)
	//}
	//return product, nil



	//product.Name


	//fmt.Println(rows)
	//defer rows.Close()
	//
	////var product Product
	////product.Name = "Update"
	////err = rows.Scan(&product.Name,&product.Stock,&product.Price)
	////if err != nil {
	////	return Product{}, nil
	////}
	////
	//////for rows.Next(){
	//////	var p Product
	//////	rows.Scan(&p.Name,&p.Stock,&p.Price)
	//////	product.Name = p.Name
	//////	product.Stock = p.Stock
	//////	product.Price = p.Price
	//////}
	//
	//var product []Product
	//for rows.Next(){
	//	var p Product
	//	p.Name = "update"
	//	err := rows.Scan(&p.ID,&p.Name,&p.Stock,&p.Price)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	product = append(product, p)
	//}
	//if err = rows.Err(); err != nil {
	//	return nil, err
	//}
	//return product, nil
}

func (d *Dependency) ProductUpdate(ctx echo.Context) error {
	productId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}
	data, err := d.RepoProductUpdate(ctx.Request().Context(),productId)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK, "OK with no transaction process updated success !!",data))
}