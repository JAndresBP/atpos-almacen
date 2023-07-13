package infrastructura

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SyncService struct {
	conn amqp.Connection
	ch   amqp.Channel
	q    amqp.Queue
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var SyncServiceInstance *SyncService

func GetSyncService(connectionString string) *SyncService {
	if SyncServiceInstance == nil {
		conn, err := amqp.Dial(connectionString)
		failOnError(err, "Failed to connect to RabbitMQ")

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")

		q, err := ch.QueueDeclare("sync", true, false, false, false, nil)
		failOnError(err, "Failed to declare a queue")

		SyncServiceInstance = &SyncService{
			conn: *conn,
			ch:   *ch,
			q:    q,
		}
	}
	return SyncServiceInstance
}

func (ss *SyncService) Publish(message []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ss.ch.PublishWithContext(ctx,
		"",
		ss.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	failOnError(err, "Failed to publish a message")
}

func (ss *SyncService) Close() {
	ss.ch.Close()
	ss.conn.Close()
}
