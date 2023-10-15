package handlers

import (
	"errors"
	"log"
	"time"

	"github.com/p2064/adder/proto"
	"github.com/p2064/pkg/db"
	"github.com/p2064/pkg/logs"
	kafka "github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) AddToEvent(ctx context.Context, in *proto.AddToEventRequest) (
	*proto.AddToEventResponse,
	error,
) {
	log.Printf("Receive message body from client: %s", in.String())
	data := db.UserEvent{
		UserID:  in.UserId,
		EventID: in.EventId,
	}

	res := db.DB.Create(&data)
	if res.Error != nil {
		return &proto.AddToEventResponse{Status: 400, Error: errors.New("Event not changed").Error()}, errors.New("Event not changed")
	}

	topic := "notify"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		logs.ErrorLogger.Fatal("failed to dial leader:", err)
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(in.String())},
	)
	if err != nil {
		logs.ErrorLogger.Fatal("failed to write messages:", err)
	}
	if err := conn.Close(); err != nil {
		logs.ErrorLogger.Fatal("failed to close writer:", err)
	}

	return &proto.AddToEventResponse{Status: 200, Error: "No error"}, nil
}
