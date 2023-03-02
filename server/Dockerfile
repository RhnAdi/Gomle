FROM golang:1.19.2-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o bin/app cmd/api/main.go

CMD "./bin/app"
