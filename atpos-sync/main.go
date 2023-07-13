package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	rabbitmqserver := os.Getenv("RABBITMQ_SERVER")
	atposcentral := os.Getenv("ATPOS_CENTRAL")

	conn, err := amqp.Dial("amqp://guest:guest@" + rabbitmqserver)

	failOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channerl")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"sync",
		true,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			data := bytes.NewBuffer(d.Body)
			response, err := http.Post("http://"+atposcentral+"/Sales", "application/json", data)

			if err != nil {
				log.Printf("%s", err)
				d.Reject(true)
				time.Sleep(1 * time.Second)
				continue
			}

			requestDump, err := httputil.DumpRequest(response.Request, true)
			if err != nil {
				log.Println(err)
			}
			log.Println(string(requestDump))

			responseDump, err := httputil.DumpResponse(response, true)
			if err != nil {
				log.Println(err)
			}
			log.Println(string(responseDump))

			if response.StatusCode == http.StatusOK {
				d.Ack(false)
				time.Sleep(1 * time.Second)
				continue
			}

			d.Reject(true)
			time.Sleep(1 * time.Second)

		}
	}()

	log.Print(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}
