package trades

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

var (
	HOST = "127.0.0.1"
	PORT = "9092"
)

func LoadHostAndPort(host string, port string) {
	HOST = "127.0.0.1"
	PORT = "9092"
}

func Publish(t string, message kafka.Message, topic string) error {
	messages := []kafka.Message{
		message,
	}
	w := kafka.Writer{
		Addr:                   kafka.TCP(HOST + ":" + PORT),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}
	defer w.Close()
	err := w.WriteMessages(context.Background(), messages...)
	if err != nil {
		log.Println("Error writing msg to kafka: ", err.Error())
		return err
	}
	log.Println("Publish msg to kafka topic: ", topic)
	return nil
}
