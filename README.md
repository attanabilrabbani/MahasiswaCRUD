# MahasiswaCRUD


CRUD with Go

How to Run:
- Create a new database in postgres
- Create a .env file containing:
PORT=4000
DB_URL="user={user} password={your password} dbname={your db name} port=5432 sslmode=disable"
- run with go run main.go
- go version: 1.22.2

Test the CRUD:
- open Postman
- import MahasiswaCRUD_postman.json as a collection
- All requests are contained within the collection
- default body for create/update (POST and PUT requests):
  {
    "Nama": {your name},
    "NPM": {your npm},
    "Semester": {your semester},
    "Ipk": {your ipk}

}

- Nama, NPM: String
- Semester: int
- Ipk: float
