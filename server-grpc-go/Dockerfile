FROM golang:1.19.0-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /server-grpc-go

CMD ["/server-grpc-go"]
