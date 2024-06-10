package main

import (
	"discount-module/model"
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	// Read json file

	data, err := ioutil.ReadFile("filename")
	if err != nil {
		log.Fatalf("Error readig Json file: %v", err)
	}

	// Parse json data

	var cart model.Cart
	if err := json.Unmarshal(data, &cart); err != nil {
		log.Fatal("Error parsing JSON data: %v", err)
	}

	// Calculate total price
	totalPrice := 0.0
	for _, item := range cart.Items {
		totalPrice += item.Price
	}
}
