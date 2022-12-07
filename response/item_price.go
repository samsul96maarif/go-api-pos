/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2022.
 */

package response

import "samsul96maarif/github.com/go-api-app/model"

type GetItemPricePaginateResponse struct {
	ItemPrices []model.ItemPrice `json:"item_prices"`
	Total      int64             `json:"total"`
	Page       uint              `json:"page"`
	Limit      uint              `json:"limit"`
}
