/*
 * Author: Samsul Ma'arfif<samsulma828@gmail.com>
 * Copyright (c) 2022.
 */

package model

type ItemPrice struct {
	BaseModel
	Name        string `json:"name"`
	Code string `json"code"`
	Price uint `json:"price"`
	CreatedBy   *uint  `json:"created_by"`
}
