package checkcampaign

import (
	"discount-module/model"
	"errors"
	"fmt"
)

func ValidateCart(cart model.Cart) error {
	for _, item := range cart.Items {
		if item.Price < 0 {
			return errors.New("item price cannot be negative")
		}
		if item.Name == "" || item.Category == "" {
			return errors.New("item name and category cannot be empty")
		}
	}

	for _, discount := range cart.Discounts {
		switch discount.Type {
		case "fixed_amount_coupon", "percentage_coupon", "category_percentage", "points_discount", "seasonal":
			// Valid types
		default:
			return fmt.Errorf("invalid discount type: %s", discount.Type)
		}

		if discount.Type == "fixed_amount_coupon" && discount.Amount < 0 {
			return errors.New("fixed amount discount cannot be negative")
		}
		if discount.Type == "percentage_coupon" && (discount.Percentage < 0 || discount.Percentage > 100) {
			return errors.New("percentage discount must be between 0 and 100")
		}
		if discount.Type == "category_percentage" && (discount.Percentage < 0 || discount.Percentage > 100) {
			return errors.New("category percentage discount must be between 0 and 100")
		}
		if discount.Type == "points_discount" && discount.Points < 0 {
			return errors.New("points discount cannot be negative")
		}
		if discount.Type == "seasonal" && (discount.EveryXTHB <= 0 || discount.DiscountYTHB < 0) {
			return errors.New("seasonal discount values must be positive")
		}
	}

	return nil
}
