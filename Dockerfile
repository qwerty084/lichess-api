FROM golang:1.22.0

WORKDIR /app

COPY ./src .

RUN go build -o app

ENTRYPOINT ["/app/app"]

EXPOSE 8080
