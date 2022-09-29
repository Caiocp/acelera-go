package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/caiocp/go-intensivo/internal/order/infra/database"
	"github.com/caiocp/go-intensivo/internal/order/usecases"
	"github.com/caiocp/go-intensivo/pkg/rabbitmq"
	_ "github.com/go-sql-driver/mysql"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	maxWorkers := 3
	wg := sync.WaitGroup{}
	
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := database.NewOrderRepository(db)

	createOrderUseCase := usecases.NewCalculateFinalPriceUseCase(repository)

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, out)

	wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		defer wg.Done()
		go worker(out, createOrderUseCase, i)
	}
	wg.Wait()
}

func worker(deliveryMessage <-chan amqp.Delivery, usecase *usecases.CalculateFinalPriceUseCase, worderId int) {
	for message := range deliveryMessage {
		var input usecases.CreateOrderInputDto
		err := json.Unmarshal(message.Body, &input)
		if err != nil {
			fmt.Println("Error unmarshalling message: ", err)
		}

		input.Tax = 10.0

		_, err = usecase.Execute(input)
		if err != nil {
			fmt.Println("Error unmarshalling message: ", err)
		}

		message.Ack(false)
		fmt.Println("Worker ", worderId, " processed message: ", input.ID)
	}
}