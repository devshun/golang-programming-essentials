FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go mod download

CMD ["go", "run", "main.go"]