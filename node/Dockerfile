FROM golang:1.23.2-alpine3.20

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

EXPOSE 8081

CMD ["go", "run", "main.go"]