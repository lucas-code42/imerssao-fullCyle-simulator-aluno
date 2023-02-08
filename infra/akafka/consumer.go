package akafka

import (
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

// NewKafkaConsumer como se fosse um construtor
func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{MsgChan: msgChan}
}

func (k *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootStrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatal("error to connect kafka -->", err.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopics")}
	c.SubscribeTopics(topics, nil)

	log.Println("Kafka consumer has been start")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg // envia a msg consumida para a variavel de msg dentro da struct (MsgChan chan ckafka.Message)
		}
	}
}
