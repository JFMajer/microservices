#!/bin/bash

export DATA_SOURCE_URL="root:verysecretpass@tcp(127.0.0.1:3306)/order"
export APPLICATION_PORT=3000
export ENV=development

go run cmd/main.go

