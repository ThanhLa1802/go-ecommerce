package initialize

import (
	"go-ecommerce-backend-api/global"
	"log"

	"github.com/segmentio/kafka-go"
)

var KafkaProducer *kafka.Writer

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Topic:    "my-topic",
		Balancer: &kafka.LeastBytes{},
	}
}

func CloseKafka() {
	if global.KafkaProducer != nil {
		if err := global.KafkaProducer.Close(); err != nil {
			log.Fatal("Failed to close Kafka producer:", err)
		}
	}
}
