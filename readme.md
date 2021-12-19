# Sample Products Restfull API

* Url 
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


# Notes

- Config.bat isn't yet ready. 
- So, I'm using this <a href="https://github.com/joho/godotenv">Go Dotenv</a> library for loading the environment
