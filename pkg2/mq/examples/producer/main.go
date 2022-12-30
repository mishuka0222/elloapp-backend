package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	kafka "gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/mq"
)

func main() {
	producer := kafka.MustKafkaProducer(&kafka.KafkaProducerConf{
		Topic:   "elloapp-test-topic",
		Brokers: []string{"127.0.0.1:9092"},
	})

	rt2 := time.Now()
	for i := 0; i < 100000; i++ {
		// rt := time.Now()
		producer.SendMessage(context.Background(), "11", []byte(fmt.Sprintf("msg: %d", rand.Int())))
		// fmt.Println("cost: ", time.Since(rt))
	}
	fmt.Println("cost: ", time.Since(rt2))
}
