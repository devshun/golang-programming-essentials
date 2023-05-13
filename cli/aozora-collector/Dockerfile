FROM golang:1.20-alpine

# 開発用
RUN apk add build-base

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go mod download

CMD ["go", "run", "main.go"]