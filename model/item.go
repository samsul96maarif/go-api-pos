/*
 * Author: Samsul Ma'arfif<samsulma828@gmail.com>
 * Copyright (c) 2022.
 */

package model

type Item struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   *uint  `json:"created_by"`
	Qty         uint   `json:"qty"`
}

type ItemWithDefaultPrice struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   *uint  `json:"created_by"`
	Qty         uint   `json:"qty"`
	Price       uint   `json:"price"`
}
