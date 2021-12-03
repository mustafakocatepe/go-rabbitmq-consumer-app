package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mustafakocatepe/go-rabbitmq-consumer-app/rabbit"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// RabbitMQ
	rc := rabbit.RabbitMQConfig{
		Schema:   "amqp",
		Username: os.Getenv("RABBITMQ_USERNAME"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
		Host:     os.Getenv("RABBITMQ_HOST"),
		Port:     os.Getenv("RABBITMQ_PORT"),
		//VHost:          "my_vhost",
		//ConnectionName: "my_app_name",
	}
	rbt := rabbit.NewRabbit(rc)

	if err := rbt.Connect(); err != nil {
		log.Fatalln("unable to connect to rabbit", err)
	}

	fmt.Println("RabbitMQ connected")

	// Consumer
	cc := rabbit.ConsumerConfig{
		ExchangeName:  "MailDirectExchange",
		ExchangeType:  "direct",
		RoutingKey:    "mail-route",
		QueueName:     "queue-mail",
		ConsumerName:  "my_app_name",
		ConsumerCount: 1,
		PrefetchCount: 1,
	}
	cc.Reconnect.MaxAttempt = 60
	cc.Reconnect.Interval = 1 * time.Second
	csm := rabbit.NewConsumer(cc, rbt)
	if err := csm.Start(); err != nil {
		log.Fatalln("unable to start consumer", err)
	}

	select {}

}
