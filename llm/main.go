package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		handleErr("Expected 'list_orders' or 'create_order' or 'take_order' subcommands", 1)
	}

	switch os.Args[1] {
	case "list_orders":
		//handle list
		if len(os.Args) != 2 {
			handleErr("No argument is needed for list_order command", 1)
		}
		handleList()
	case "create_order":
		//handle create
		if len(os.Args[2:]) != 2 {
			handleErr("Require the initial location and/or destination", 1)
		}
		fmt.Println(handleCreate(os.Args[2], os.Args[3]))
	case "take_order":
		//handle take
		if len(os.Args[2:]) != 1 {
			handleErr("Wrong input for 'take_order' commmand", 1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			handleErr("Non-numeric order id", 1)
		}
		handleTake(id)
	default:
		//handle err
		handleErr("Unknow subcommand", 1)
	}

}

func handleErr(errMsg string, errCode int) {
	fmt.Println("ERR: " + errMsg)
	os.Exit(errCode)
}

func handleList() {
	orders := getOrders()
	for _, order := range orders {
		if !order.Taken {
			fmt.Printf("%v,%v,%v\n", order.Id, order.From, order.To)
		}
	}
	return
}

func handleCreate(from string, to string) (id int) {
	if from == "" || to == "" {
		handleErr("All fields are required for creating a order", 1)
	}
	return makeOrder(from, to)
}

func handleTake(id int) {
	if id > len(getOrders()) {
		handleErr("Order does not exist", 1)
	}
	if !takeOrder(id) {
		handleErr("Order already taken", 1)
	}
	return
}
