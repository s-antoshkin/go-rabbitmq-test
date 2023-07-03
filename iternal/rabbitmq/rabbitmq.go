package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Service interface {
	Connect() error
	Publish(message string) error
	Consume()
}

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func (r *RabbitMQ) Connect() error {
	fmt.Println("Connecting to RabbitMQ...")

	var err error

	r.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}

	fmt.Println("Successfully connecting to RabbitMQ")

	r.Channel, err = r.Conn.Channel()
	if err != nil {
		return err
	}

	_, err = r.Channel.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQ) Publish(msg string) error {
	err := r.Channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)

	if err != nil {
		return err
	}

	fmt.Println("Successfully published message to queue")

	return nil
}

func (r *RabbitMQ) Consume() {
	msgs, err := r.Channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	for msg := range msgs {
		fmt.Printf("Received message: %s\n", msg.Body)
	}
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
