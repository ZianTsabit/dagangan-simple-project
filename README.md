# Dagangan Simple Project

Simple Project For Dagangan Backend Developer Internship

## Description

Using Go, Build a REST API for Product entity with following feature:

* Create, Read, Update, Delete
* Filtering
* Pagination

## Requirements & Tech Stack

1. Docker
2. Docker Compose
3. PostgreSQL
4. Go
5. GO

## How to Run and Build?

```shell
sudo docker-compose up --build
```

## Documentation

1. Create New Product
   **POST** `http://localhost:8000/product`
   contoh body yang digunakan untuk melakukan request

   ```shell
   {
   	"name": "Macbook Air",
   	"price": 25000000,
   	"quantity": 3
   }
   ```
2. Read Product
   **GET** `http://localhost:8000/product`

   - pagination (default : page_size=5)
     **GET** `http://localhost:8000/product?page_size=5&page_num=1`
   - filtering
     **GET** `http://localhost:8000/product?name=Macbook`

     **GET** `http://localhost:8000/product?price=25000000`

     **GET** `http://localhost:8000/product?name=Macbook&price=25000000`
3. Update Product
   **PUT** `http://localhost:8000/product/:id`

   ```shell
   {
   	"name": "Macbook Air M2 2023",
   	"price": 27000000,
   	"quantity": 4
   }
   ```
4. Delete Product
   **DELETE** `http://localhost:8000/product/:id`
