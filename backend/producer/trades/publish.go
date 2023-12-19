package trades

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

var (
	HOST string
	PORT string
)

func LoadHostAndPort(host string, port string) {
	HOST = host
	PORT = port
}

func Publish(t string, message kafka.Message, topic string) error {

	messages := []kafka.Message{
		message,
	}

	w := kafka.Writer{
		Addr:                   kafka.TCP(HOST + ":" + PORT), //127.0.0.1:9092 or kafka:9092 in docker
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}
	defer w.Close()

	err := w.WriteMessages(context.Background(), messages...)
	if err != nil {
		log.Println("Error writing messages to Kafka: ", err.Error())
		return err
	}

	log.Println("Publish messages to Kafka on topic: ", topic)

	return nil
}
