package main

import (
	"database/sql"

	"github.com/caiocp/go-intensivo/internal/order/infra/database"
	"github.com/caiocp/go-intensivo/internal/order/usecases"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := database.NewOrderRepository(db)

	createUseCase := usecases.NewCalculateFinalPriceUseCase(repository)

	input := usecases.CreateOrderInputDto{
		ID:    "123",
		Price: 100,
		Tax:   10,
	}

	output, err := createUseCase.Execute(input)
	if err != nil {
		panic(err)
	}

	println(output.FinalPrice)
}