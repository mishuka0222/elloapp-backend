package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"
)

func main() {
	job := kafka.MustKafkaConsumer(&kafka.KafkaConsumerConf{
		Topics:  []string{"elloapp-test-topic"},
		Brokers: []string{"127.0.0.1:9092"},
		Group:   "elloapp-test-group-job",
	})

	job.RegisterHandlers("elloapp-test-topic",
		func(ctx context.Context, key string, value []byte) {
			fmt.Println("key: ", key, ", value: ", string(value))
		})

	defer job.Stop()
	go job.Start()
	// signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Println("get a signal ", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			// job.Close()
			fmt.Println("exit...")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
