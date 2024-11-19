package pkg

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
)

type Handler struct {
	r IRepository
}

func NewHandler(r IRepository) *Handler {
	return &Handler{r: r}
}

func (h *Handler) Insert(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(`{"message": "success"}`))
}

func (h *Handler) OnMessageFromOrderService() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:29092"},
		Topic:   "inventory",
		GroupID: "group1",
	})
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("failed to read message: %s\n", err)
		}
		if string(msg.Value) == "ORDER CREATED" {
			number, err := h.r.Get(context.Background())
			if err != nil || number == 0 {
				err = ProduceMessage("localhost:29092", "order", "RUN OUT")
				if err != nil {
					log.Fatalf("failed to produce message: %v", err)
				}
			} else {
				err = ProduceMessage("localhost:29092", "order", "AVAILABLE")
				if err != nil {
					log.Fatalf("failed to produce message: %v", err)
				}
			}
		}
	}
}
