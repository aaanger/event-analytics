// Test gRPC client for generating random events

package grpc

import (
	"context"
	pb "github.com/aaanger/event-analytics/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to grpc server: %s", err)
	}

	defer conn.Close()

	client := pb.NewEventClient(conn)

	for i := 0; i < 100; i++ {
		req := generateEvent()
		sendEvent(client, req)
		time.Sleep(time.Second)
	}
}

func generateEvent() *pb.EventRequest {
	urls := []string{
		"/home",
		"/contacts",
		"/cart",
		"/products/1",
		"/products/2",
		"/orders/create",
		"/orders/cancel",
	}

	return &pb.EventRequest{
		UserId:    int64(rand.Intn(10000)),
		Url:       urls[rand.Intn(len(urls))],
		Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
	}
}

func sendEvent(client pb.EventClient, req *pb.EventRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.SendEvent(ctx, req)
	if err != nil {
		log.Printf("failed to send event: %s", err)
	} else {
		log.Printf("event sent: %+v", req)
	}
}
