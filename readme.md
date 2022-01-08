# Sample Products Restfull API
```bash
You guys must be run mysql and pgsql at the same time.
Because, i'm using mysql and pgsql in this project.
In pgsql, u must setup the name=postgres, 
dbname=postgres, password=postgres.
But, literally you can custom
by yourself. Happy Coding.
```
# Using Postgres-Sql Database

## Get All

* Url
  - `GET` Method
```http request
http://localhost:8000/api/users
```

body response
```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "id": "73e12106-207e-4693-9c0d-3147d6ab606a",
      "first_name": "Wahidin",
      "last_name": "Aji",
      "email": "a17wahidin@gmail.com",
      "phone_number": 123456789012
    },
    {
      "id": "44c22cb3-ff6c-4043-8c79-8a5506ce11e9",
      "first_name": "Tia",
      "last_name": "Ulul Putri",
      "email": "tiaulul.putri@mail.com",
      "phone_number": 123123123123
    },
    {
      "id": "d90f8110-039a-47f4-a164-37d807f77ab5",
      "first_name": "Omoy",
      "last_name": "Bungsu Putri",
      "email": "omoy.putri@mail.com",
      "phone_number": 987654321098
    }
  ]
}
```

## Get By UUID
* Url
  - `GET` Method
```http request
http://localhost:8000/api/users/d90f8110-039a-47f4-a164-37d807f77ab5
```
params uuid request
```bash
uuid <- d90f8110-039a-47f4-a164-37d807f77ab5
```
body response

- data found
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "id": "d90f8110-039a-47f4-a164-37d807f77ab5",
    "first_name": "Omoy",
    "last_name": "Bungsu Putri",
    "email": "omoy.putri@mail.com",
    "phone_number": 987654321098
  }
}
```
- data not found
```json
{
  "code": 404,
  "status": "Not Found",
  "data": "user id was not found : %!(EXTRA <nil>)"
}
```

## Update By UUID
* Url
  - `PATCH` Method
```http request
http://localhost:8000/api/users/d90f8110-039a-47f4-a164-37d807f77ab5
```
body request
```json
{
    "first_name" : "Omoyi",
    "last_name" : "Bungsu Putri",
    "email" : "omoy.putri@mail.com",
    "phone_number" : 987654321098
}
```
body response
- data found, then update
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "id": "d90f8110-039a-47f4-a164-37d807f77ab5",
    "first_name": "Omoy update",
    "last_name": "Bungsu Putri",
    "email": "omoy.putri@mail.com",
    "phone_number": 987654321098
  }
}
```
- data not found
```json
{
  "code": 404,
  "status": "Not Found",
  "data": "user id was not found : %!(EXTRA <nil>)"
}
```


# Using Mysql Database

## Get All

* Url
    - `GET` Method             
```http request
http://localhost:8000/products
```

body response
```json
[
  {
    "id": 1,
    "name": "Product One",
    "stock": 15,
    "price": 10000.01
  },
  {
    "id": 2,
    "name": "Product Two",
    "stock": 20,
    "price": 20000.02
  },
  {
    "id": 3,
    "name": "Product Three",
    "stock": 25,
    "price": 25000.03
  }
]
```

## Get By Id
* Url
  - `GET` Method
```http request
http://localhost:8000/products/3
```

body response

 - data found
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "id": 3,
    "name": "Product Three",
    "stock": 25,
    "price": 25000.03
  }
}
```
  - data not found
```json
{
  "code": 200,
  "status": "OK",
  "data": null
}
```

## Update by id
* Url
  - `PATCH` Method
```http request
http://localhost:8000/products/3
```

body request
```json
{
    "name": "update 3",
    "stock":1002,
    "price":10.99
}
```
body response
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "id": 3,
    "name": "update 3",
    "stock": 1002,
    "price": 10.99
  }
}
```


# Notes

- Config.bat isn't yet ready. 
- So, I'm using this <a href="https://github.com/joho/godotenv">Go Dotenv</a> library for loading the environment
