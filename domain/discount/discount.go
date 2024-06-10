package discount

import (
	filtertype "discount-module/domain/filterType"
	"discount-module/model"
	"math"
)

func ApplyDiscounts(totalPrice float64, items []model.Item, discounts []model.Discount) float64 {
	// filter and apply coupon
	couponDiscount := filtertype.FilterCouponDiscount(discounts)
	totalPrice = ApplyCouponDiscounts(totalPrice, couponDiscount)

	//filter and apply On Top
	onTopDiscounts := filtertype.FilterOnTopDiscount(discounts)
	totalPrice = ApplyOnTopDiscounts(totalPrice, items, onTopDiscounts)

	seasonalDiscounts := filtertype.FilterSeasonalDiscounts(discounts)
	totalPrice = ApplySeasonalDiscounts(totalPrice, seasonalDiscounts)
	return totalPrice
}

func ApplyCouponDiscounts(totalPrice float64, discounts []model.Discount) float64 {
	for _, discount := range discounts {
		if discount.Type == "fixed_amount_coupon" {
			totalPrice -= discount.Amount
		} else if discount.Type == "percentage_coupon" {
			totalPrice -= totalPrice * discount.Percentage / 100
		}
	}
	return totalPrice
}

func ApplyOnTopDiscounts(totalPrice float64, items []model.Item, discounts []model.Discount) float64 {
	for _, discount := range discounts {
		if discount.Type == "category_percentage" {
			for _, item := range items {
				if item.Category == discount.Category {
					totalPrice -= item.Price * discount.Percentage / 100
				}
			}
		} else if discount.Type == "point_discount" {
			pointsDiscount := float64(discount.Points)
			maxDiscount := totalPrice * 0.2
			if pointsDiscount > maxDiscount {
				pointsDiscount = maxDiscount
			}
			totalPrice -= pointsDiscount
		}
	}
	return totalPrice
}

func ApplySeasonalDiscounts(totalPrice float64, discounts []model.Discount) float64 {
	for _, discount := range discounts {
		if discount.Type == "seasonal" {
			discountMultiplier := math.Floor(totalPrice / discount.EveryXTHB)
			totalPrice -= discountMultiplier * discount.DiscountYTHB
		}
	}
	return totalPrice
}
