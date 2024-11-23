package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
  configMap := &kafka.ConfigMap{
    "bootstrap.servers": "full-cycle-apache-kafka-kafka-1:9092",
    "client.id": "goapp",
    "group.id": "goapp",
  }
  consumer, err := kafka.NewConsumer(configMap)
  if err != nil {
    fmt.Println("Error creating new consumer", err.Error())
  }
  topics := []string{"teste"}
  consumer.SubscribeTopics(topics, nil)
  for {
    message, err := consumer.ReadMessage(-1)
    if err == nil {
      fmt.Println(string(message.Value), message.TopicPartition)
    }
  }
}
