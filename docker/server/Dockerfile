FROM golang:1.17.5-bullseye

RUN go install github.com/cosmtrek/air@v1.29.0

WORKDIR /go/src/

COPY go.mod go.sum ./
RUN go mod download
