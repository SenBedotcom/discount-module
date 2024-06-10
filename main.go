package main

import (
	"discount-module/domain/discount"
	"discount-module/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Read json file

	data, err := ioutil.ReadFile("cart.json")
	if err != nil {
		log.Fatalf("Error readig Json file: %v", err)
	}

	// Parse json data

	var cart model.Cart
	if err := json.Unmarshal(data, &cart); err != nil {
		log.Fatalf("Error parsing JSON data: %v", err)
	}

	// Calculate total price
	totalPrice := 0.0
	for _, item := range cart.Items {
		totalPrice += item.Price
	}

	finalPrice := discount.ApplyDiscounts(totalPrice, cart.Items, cart.Discounts)
	fmt.Printf("Final Price: %.2f THB\n", finalPrice)
}
