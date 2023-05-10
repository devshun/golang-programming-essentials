FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "run", "main.go"]