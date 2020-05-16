FROM golang:1.12.10-alpine

WORKDIR /terpusat

COPY . .

RUN apk update && apk upgrade; \
    apk add --no-cache bash git openssh; \
    go mod tidy; \
    rm -rf clone checkout fetch pull push types utils main.go Makefile;

