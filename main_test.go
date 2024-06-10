// package main

// import (
// 	checkcampaign "discount-module/domain/checkCampaign"
// 	"discount-module/domain/discount"
// 	"discount-module/model"
// 	"encoding/json"
// 	"testing"
// )

// func TestApplyDiscounts(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		cartJSON string
// 		expected float64
// 	}{
// 		{
// 			name: "Fixed Amount Coupon",
// 			cartJSON: `{
// 				"items": [
// 					{"name": "T-Shirt", "category": "Clothing", "price": 350},
// 					{"name": "Hat", "category": "Accessories", "price": 250}
// 				],
// 				"discounts": [
// 					{"type": "fixed_amount_coupon", "amount": 50}
// 				]
// 			}`,
// 			expected: 550 - 50,
// 		},
// 		{
// 			name: "Percentage Coupon",
// 			cartJSON: `{
// 				"items": [
// 					{"name": "T-Shirt", "category": "Clothing", "price": 350},
// 					{"name": "Hat", "category": "Accessories", "price": 250}
// 				],
// 				"discounts": [
// 					{"type": "percentage_coupon", "percentage": 10}
// 				]
// 			}`,
// 			expected: 600 - 600*0.1,
// 		},
// 		{
// 			name: "On Top Discount",
// 			cartJSON: `{
// 				"items": [
// 					{"name": "T-Shirt", "category": "Clothing", "price": 350},
// 					{"name": "Hoodie", "category": "Clothing", "price": 700},
// 					{"name": "Watch", "category": "Accessories", "price": 850},
// 					{"name": "Bag", "category": "Accessories", "price": 640}
// 				],
// 				"discounts": [
// 					{"type": "category_percentage", "category": "Clothing", "percentage": 15}
// 				]
// 			}`,
// 			expected: 2540 - (350*0.15 + 700*0.15),
// 		},
// 		{
// 			name: "Points Discount",
// 			cartJSON: `{
// 				"items": [
// 					{"name": "T-Shirt", "category": "Clothing", "price": 350},
// 					{"name": "Hat", "category": "Accessories", "price": 250},
// 					{"name": "Belt", "category": "Accessories", "price": 230}
// 				],
// 				"discounts": [
// 					{"type": "points_discount", "points": 68}
// 				]
// 			}`,
// 			expected: 830 - 68,
// 		},
// 		{
// 			name: "Seasonal Discount",
// 			cartJSON: `{
// 				"items": [
// 					{"name": "T-Shirt", "category": "Clothing", "price": 350},
// 					{"name": "Hat", "category": "Accessories", "price": 250},
// 					{"name": "Belt", "category": "Accessories", "price": 230}
// 				],
// 				"discounts": [
// 					{"type": "seasonal", "every_x_thb": 300, "discount_y_thb": 40}
// 				]
// 			}`,
// 			expected: 830 - 80,
// 		},
// 		{
// 			name: "Multiple Discounts",
// 			cartJSON: `{
// 				"items": [
// 					{"name": "T-Shirt", "category": "Clothing", "price": 350},
// 					{"name": "Hat", "category": "Accessories", "price": 250},
// 					{"name": "Belt", "category": "Accessories", "price": 230}
// 				],
// 				"discounts": [
// 					{"type": "fixed_amount_coupon", "amount": 50},
// 					{"type": "category_percentage", "category": "Clothing", "percentage": 10},
// 					{"type": "seasonal", "every_x_thb": 300, "discount_y_thb": 40}
// 				]
// 			}`,
// 			expected: 830 - 50 - 35 - 80,
// 		},
// 		{
// 			name: "Negative Item Price",
// 			cartJSON: `{
// 				"items": [
// 					{"name": "T-Shirt", "category": "Clothing", "price": -350},
// 					{"name": "Hat", "category": "Accessories", "price": 250}
// 				],
// 				"discounts": []
// 			}`,
// 			expected: -1, // Invalid cart should be handled and not calculate a price
// 		},
// 		{
// 			name: "Invalid Discount Type",
// 			cartJSON: `{
// 				"items": [
// 					{"name": "T-Shirt", "category": "Clothing", "price": 350},
// 					{"name": "Hat", "category": "Accessories", "price": 250}
// 				],
// 				"discounts": [
// 					{"type": "invalid_type", "amount": 50}
// 				]
// 			}`,
// 			expected: -1, // Invalid discount should be handled and not apply any discount
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var cart model.Cart
// 			if err := json.Unmarshal([]byte(tt.cartJSON), &cart); err != nil {
// 				t.Fatalf("Error parsing JSON: %v", err)
// 			}

// 			err := checkcampaign.ValidateCart(cart)
// 			if tt.expected == -1 {
// 				if err == nil {
// 					t.Errorf("Expected validation error, but got none")
// 				}
// 				return
// 			}

// 			if err != nil {
// 				t.Fatalf("Unexpected validation error: %v", err)
// 			}

// 			totalPrice := 0.0
// 			for _, item := range cart.Items {
// 				totalPrice += item.Price
// 			}

// 			finalPrice := discount.ApplyDiscounts(totalPrice, cart.Items, cart.Discounts)
// 			if finalPrice != tt.expected {
// 				t.Errorf("Expected final price to be %.2f, but got %.2f", tt.expected, finalPrice)
// 			}
// 		})
// 	}
// }

package main

import (
	checkcampaign "discount-module/domain/checkCampaign"
	"discount-module/domain/discount"
	"discount-module/model"
	"encoding/json"
	"testing"
)

func TestApplyDiscounts(t *testing.T) {
	tests := []struct {
		name     string
		cartJSON string
		expected float64
	}{
		{
			name: "Fixed Amount Coupon",
			cartJSON: `{
				"items": [
					{"name": "T-Shirt", "category": "Clothing", "price": 350},
					{"name": "Hat", "category": "Accessories", "price": 250}
				],
				"discounts": [
					{"type": "fixed_amount_coupon", "amount": 50}
				]
			}`,
			expected: 600 - 50,
		},
		{
			name: "Percentage Coupon",
			cartJSON: `{
				"items": [
					{"name": "T-Shirt", "category": "Clothing", "price": 350},
					{"name": "Hat", "category": "Accessories", "price": 250}
				],
				"discounts": [
					{"type": "percentage_coupon", "percentage": 10}
				]
			}`,
			expected: 600 - 600*0.1,
		},
		{
			name: "On Top Discount",
			cartJSON: `{
				"items": [
					{"name": "T-Shirt", "category": "Clothing", "price": 350},
					{"name": "Hoodie", "category": "Clothing", "price": 700},
					{"name": "Watch", "category": "Accessories", "price": 850},
					{"name": "Bag", "category": "Accessories", "price": 640}
				],
				"discounts": [
					{"type": "category_percentage", "category": "Clothing", "percentage": 15}
				]
			}`,
			expected: 2540 - (350*0.15 + 700*0.15),
		},
		{
			name: "Points Discount",
			cartJSON: `{
				"items": [
					{"name": "T-Shirt", "category": "Clothing", "price": 350},
					{"name": "Hat", "category": "Accessories", "price": 250},
					{"name": "Belt", "category": "Accessories", "price": 230}
				],
				"discounts": [
					{"type": "points_discount", "points": 68}
				]
			}`,
			expected: 830 - 68,
		},
		{
			name: "Seasonal Discount",
			cartJSON: `{
				"items": [
					{"name": "T-Shirt", "category": "Clothing", "price": 350},
					{"name": "Hat", "category": "Accessories", "price": 250},
					{"name": "Belt", "category": "Accessories", "price": 230}
				],
				"discounts": [
					{"type": "seasonal", "every_x_thb": 300, "discount_y_thb": 40}
				]
			}`,
			expected: 830 - 80,
		},
		{
			name: "Multiple Discounts",
			cartJSON: `{
				"items": [
					{"name": "T-Shirt", "category": "Clothing", "price": 350},
					{"name": "Hat", "category": "Accessories", "price": 250},
					{"name": "Belt", "category": "Accessories", "price": 230}
				],
				"discounts": [
					{"type": "fixed_amount_coupon", "amount": 50},
					{"type": "category_percentage", "category": "Clothing", "percentage": 10},
					{"type": "seasonal", "every_x_thb": 300, "discount_y_thb": 40}
				]
			}`,
			expected: 830 - 50 - 35 - 80,
		},
		{
			name: "Negative Item Price",
			cartJSON: `{
				"items": [
					{"name": "T-Shirt", "category": "Clothing", "price": -350},
					{"name": "Hat", "category": "Accessories", "price": 250}
				],
				"discounts": []
			}`,
			expected: -1, // Invalid cart should be handled and not calculate a price
		},
		{
			name: "Invalid Discount Type",
			cartJSON: `{
				"items": [
					{"name": "T-Shirt", "category": "Clothing", "price": 350},
					{"name": "Hat", "category": "Accessories", "price": 250}
				],
				"discounts": [
					{"type": "invalid_type", "amount": 50}
				]
			}`,
			expected: -1, // Invalid discount should be handled and not apply any discount
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cart model.Cart
			if err := json.Unmarshal([]byte(tt.cartJSON), &cart); err != nil {
				t.Fatalf("Error parsing JSON: %v", err)
			}

			err := checkcampaign.ValidateCart(cart)
			if tt.expected == -1 {
				if err == nil {
					t.Errorf("Expected validation error, but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected validation error: %v", err)
			}

			totalPrice := 0.0
			for _, item := range cart.Items {
				totalPrice += item.Price
			}

			finalPrice := discount.ApplyDiscounts(totalPrice, cart.Items, cart.Discounts)
			if finalPrice != tt.expected {
				t.Errorf("Test %s: Expected final price to be %.2f, but got %.2f", tt.name, tt.expected, finalPrice)
			}
		})
	}
}
