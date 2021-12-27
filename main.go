package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"products-restapi/components/products"
	"time"
)

func init() {
	if appName := os.Getenv("APP_NAME"); appName == "" {
		log.Fatal("Please provide the ENVIRONMENT value of (APP_NAME) on the .env file")
	}
	if dbUser := os.Getenv("DB_USER"); dbUser == "" {
		log.Fatal("Please provide the ENVIRONMENT value of (DB_USER) on the .env file")
	}
	//enable pass if your db has password
	//if dbPass := os.Getenv("DB_PASS"); dbPass == "" {
	//	log.Fatal("Please provide the ENVIRONMENT value of (DB_PASS) on the .env file")
	//}
	if dbName := os.Getenv("DB_NAME"); dbName == "" {
		log.Fatal("Please provide the ENVIRONMENT value of (DB_NAME) on the .env file")
	}
	log.Println("Passed the environment variable check")
}

func main() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, pass, name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("during opening a postgres client:", errors.New("Invalid connection!!!"), err)
	}
	db.SetMaxIdleConns(10)                  //min 10 connection
	db.SetMaxOpenConns(100)                 //max 100 connection
	db.SetConnMaxIdleTime(5 * time.Minute)  //if iin 5 minutes nothing happen? db will close the connection
	db.SetConnMaxLifetime(60 * time.Minute) //after 60 minutes, the connection will create new connection
	defer db.Close()

	e := echo.New()
	e.Server.ReadTimeout = time.Duration(2) * time.Minute  //SERVER_READ_TIMEOUT_IN_MINUTE
	e.Server.WriteTimeout = time.Duration(2) * time.Minute //SERVER_WRITE_TIMEOUT_IN_MINUTE

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//{
	//	"name":"Aji",
	//	"email":"aji@mail.com"
	//}
	//products
	product := products.Dependency{DB: db}
	e.GET("/products", product.GetAll)
	e.GET("/products/:id", product.GetById)

	//with no transaction process
	api := e.Group("/api")
	api.GET("/products", product.ProductsAll)
	api.PUT("/products/:id", product.ProductUpdate)

	//running server
	server := new(http.Server)
	server.Addr = ":8000"
	e.Logger.Print(os.Getenv("APP_NAME"), " is running on http://localhost", server.Addr)
	e.Logger.Fatal(e.StartServer(server))
}
