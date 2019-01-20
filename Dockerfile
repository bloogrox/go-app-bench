FROM golang:1.11 as build-env

ENV GO111MODULE on

WORKDIR /go/src/app

COPY . .

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o web

EXPOSE 80

CMD ["./web"]
