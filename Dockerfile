FROM golang:1.23.4 AS grpc-server


COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o grpc-server ./cmd

EXPOSE 50051

CMD ["./grpc-server"]