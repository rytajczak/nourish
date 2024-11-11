FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /user-service

EXPOSE 8080

CMD ["/user-service"]