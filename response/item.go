/*
 * Author: Samsul Ma'arfi <samsulma828@gmail.com>
 * Copyright (c) 2022.
 */

package response

import "samsul96maarif/github.com/go-api-app/model"

type GetItemPaginateResponse struct {
	Items []model.ItemWithDefaultPrice `json:"items"`
	Total int64                        `json:"total"`
	Page  uint                         `json:"page"`
	Limit uint                         `json:"limit"`
}
