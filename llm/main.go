package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		handleErr("expected 'list_orders' or 'create_order' or 'take_order' subcommands", 1)
	}

	switch os.Args[1] {
	case "list_orders":
		//handle list
		// handleList(listCmd)
		if len(os.Args) != 2 {
			handleErr("no argument is needed for list_order command", 1)
		}
		handleList()
	case "create_order":
		//handle create
		if len(os.Args[2:]) != 2 {
			handleErr("require the initial location and/or destination", 1)
		}
		fmt.Println(handleCreate(os.Args[2], os.Args[3]))
	case "take_order":
		//handle take
		if len(os.Args[2:]) != 1 {
			handleErr("wrong input for take_order commmand", 1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			handleErr("non-numeric order id", 1)
		}
		handleTake(id)
	default:
		//handle err
		handleErr("unknow subcommand", 1)
	}

}

func handleErr(errMsg string, errCode int) {
	fmt.Println(errMsg)
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
		handleErr("all field are required for creating a order", 1)
	}
	return makeOrder(from, to)
}

func handleTake(id int) {
	if id > len(getOrders()) {
		handleErr("order does not exist", 1)
	}
	if !takeOrder(id) {
		handleErr("order already taken", 1)
	}
	return
}
