/*
 * Author: Samsul Ma'arfif<samsulma828@gmail.com>
 * Copyright (c) 2022.
 */

package model

type ItemPrice struct {
	BaseModel
	Name      string `json:"name"`
	Code      string `json:"code"`
	CreatedBy *uint  `json:"created_by"`
	ItemId    uint   `json:"item_id"`
	Price     uint   `json:"price"`
	IsDefault bool   `json:"is_default"`
}
