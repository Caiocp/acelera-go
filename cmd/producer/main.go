package main

import (
	"encoding/json"
	"math/rand"

	"github.com/caiocp/go-intensivo/pkg/rabbitmq"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Order struct {
		ID string
		Price float64
}

func GenerateOrders() Order {
	return Order{ID: uuid.New().String(), Price: rand.Float64() * 100}
}

func Notify(channel *amqp.Channel, order Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = channel.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	return err
}

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	for i:= 0; i < 100000; i++ {
		order := GenerateOrders()
		err = Notify(ch, order)
		if err != nil {
			panic(err)
		}
		// fmt.Println("Order created: ", order)
	}
}