FROM golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o grpc-server ./cmd/server

EXPOSE 50051

CMD ["./grpc-server"]