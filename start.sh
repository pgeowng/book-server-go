#!/bin/sh

export PG_URL="postgres://localhost:5432/books"
export HTTP_ADDR="localhost:1314"
go run main.go