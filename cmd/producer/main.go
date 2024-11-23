package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
) 

func main() {
  deliveryChannel := make(chan kafka.Event)
  producer := NewKafkaProducer()
  Publish("Mensagem de transferencia", "teste", producer, []byte("transferencia"), deliveryChannel)
  go DeliveryReport(deliveryChannel)
  producer.Flush(3000)  
}

func NewKafkaProducer() *kafka.Producer {
  configMap := &kafka.ConfigMap{
    "bootstrap.servers": "full-cycle-apache-kafka-kafka-1:9092",
    "delivery.timeout.ms": "0",
    "acks": "all",
    "enable.idempotence": "true",
  }
  p, err := kafka.NewProducer(configMap)
  if err != nil {
    log.Println(err.Error())
  }
  return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChannel chan kafka.Event) error {
  message := &kafka.Message{
    TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
    Value:          []byte(msg),
    Key:            key,
  }
  err := producer.Produce(message, deliveryChannel)
  if err != nil {
    return err
  }
  return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
  for e := range deliveryChan {
    switch ev := e.(type) {
    case *kafka.Message:
      if ev.TopicPartition.Error != nil {
        fmt.Println("Delivery failed", ev.TopicPartition)
      } else {
        fmt.Println("Delivered message to", ev.TopicPartition)
      }
    }
  }
}
