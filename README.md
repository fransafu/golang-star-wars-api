# golang-star-wars-api

A simple star wars API made in Golang

The goal of this project is to make a hexagonal architecture POC using the star-wars API as outgoing.

**Next steps**
* Add cache layer using **dragonflydb**

## Getting started

**Run redis**
* docker-compose up -d redis

**Run dragonflydb**
* docker-compose up -d dragonflydb

**How to run the program**
* go run cmd/api/main.go
