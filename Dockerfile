FROM golang:1.21.6

WORKDIR /app

COPY ./src .

RUN go build -o app

ENTRYPOINT ["/app/app"]

EXPOSE 8080
