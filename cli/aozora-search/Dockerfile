FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go mod init aozora-search
RUN go mod tidy

CMD ["go", "run", "main.go"]