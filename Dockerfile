FROM golang:1.21

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build  -o /validator ./cmd/validator-rest/main.go

CMD ["/validator"]