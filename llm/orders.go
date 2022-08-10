package main

import (
	"encoding/json"
	"io/ioutil"
)

type order struct {
	Id    int    `json:"id"`
	From  string `json:"from"`
	To    string `json:"to"`
	Taken bool   `json:"taken"`
}

func getOrders() (orders []order) {
	fileBytes, err := ioutil.ReadFile("./orders.json")
	if err != nil {
		handleErr("failed to read file", 1)
	}

	err = json.Unmarshal(fileBytes, &orders)
	if err != nil {
		handleErr("failed to unmarshal json", 1)
	}

	return orders
}

func makeOrder(from string, to string) (id int) {
	orders := getOrders()

	uuid := 0
	if len(orders) == 0 {
		uuid = 1
	} else {
		uuid = orders[len(orders)-1].Id + 1
	}

	order := order{
		Id:    uuid,
		From:  from,
		To:    to,
		Taken: false,
	}
	orders = append(orders, order)

	saveOrders(orders)

	return (uuid)
}

func saveOrders(orders []order) {
	orderByte, err := json.Marshal(orders)
	if err != nil {
		handleErr("failed to marshal json", 1)
	}

	err = ioutil.WriteFile("./orders.json", orderByte, 0644)

	if err != nil {
		handleErr("failed to write file", 1)
	}

	return
}

func takeOrder(id int) (success bool) {
	orders := getOrders()

	for i, order := range orders {
		if order.Id == id {
			if !order.Taken {
				orders[i].Taken = true
				saveOrders(orders)
				return true
			}
			break
		}
	}
	return false

}
