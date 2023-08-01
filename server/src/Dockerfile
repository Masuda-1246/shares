## Build
FROM golang:latest AS build

WORKDIR /go/src/github.com/Masuda-1246/shares

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

## Deploy
# FROM gcr.io/distroless/base-debian10

