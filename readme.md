# Sample Products Restfull API

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
