FROM golang:1.20-alpine

WORKDIR /app
COPY . .

RUN go build -o auth .

EXPOSE 8080
ENTRYPOINT ["./auth"]
