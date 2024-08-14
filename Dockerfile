FROM golang:1.23.0-alpine

WORKDIR /app
COPY . .

RUN go build -o app

EXPOSE 8080:8080

CMD ["./app"]