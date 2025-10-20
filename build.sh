#!/bin/bash

mkdir -p ./out

# clean old out
rm -rf ./out/*

# build app
templ generate
bunx @tailwindcss/cli -i ./src/input.css -o ./static/output.css 
go build -o ./out/server ./cmd/server

cp .env ./out/.env
cp -r src ./out/
cp -r static ./out/
