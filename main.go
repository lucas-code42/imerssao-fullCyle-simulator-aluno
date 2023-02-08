package main

import (
	"log"

	"github.com/imersao-full-cycle/simulator/infra/akafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error load .env")
	}
}

func main() {

	// msgChan := make(chan *kafka.Message)
	// consumer := akafka.NewKafkaConsumer(msgChan)
	// go consumer.Consume()

	// for msg := range msgChan{
	// 	log.Println(string(msg.Value))
	// }

	producer := akafka.NewKafkaProducer()
	akafka.Publish("ola", "readtest", producer)

	for {
		_ = 1
	}

	// route := route.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }

	// route.LoadPositions()
	// stringJson, _ := route.ExportJSONPositions()

	// for _, v := range stringJson {
	// 	fmt.Println(v)
	// }
}
