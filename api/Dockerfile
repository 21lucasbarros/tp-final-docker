FROM golang:alpine

WORKDIR /app

COPY main.go .

RUN go build -o api main.go

EXPOSE 3000

CMD ["./api"]