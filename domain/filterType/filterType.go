package filtertype

import (
	cart "discount-module/model"
)

func FilterCouponDiscount(discounts []cart.Discount) []cart.Discount {
	var selectDiscounts []cart.Discount
	fixAmountApply, percentageApply := false, false

	for _, discount := range discounts {
		if discount.Type == "fixed_amount_coupon" && !fixAmountApply && !percentageApply {
			selectDiscounts = append(selectDiscounts, discount)
			fixAmountApply = true
		} else if discount.Type == "percentage_coupon" && !fixAmountApply && !percentageApply {
			selectDiscounts = append(selectDiscounts, discount)
			percentageApply = true
		}
	}

	return selectDiscounts
}

func FilterOnTopDiscount(discounts []cart.Discount) []cart.Discount {
	var selectDiscounts []cart.Discount
	categoryApply := make(map[string]bool)

	for _, discount := range discounts {
		if discount.Type == "category_percentage" {
			if !categoryApply[discount.Category] {
				selectDiscounts = append(selectDiscounts, discount)
				categoryApply[discount.Category] = true
			}
		} else if discount.Type == "points_discount" {
			selectDiscounts = append(selectDiscounts, discount)
		}
	}
	return selectDiscounts
}

func FilterSeasonalDiscounts(discounts []cart.Discount) []cart.Discount {
	var selectedDiscounts []cart.Discount

	for _, discount := range discounts {
		if discount.Type == "seasonal" {
			selectedDiscounts = append(selectedDiscounts, discount)
		}
	}
	return selectedDiscounts
}
