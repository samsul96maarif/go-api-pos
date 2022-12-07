/*
 * Author: Samsul Ma'arfi <samsulma828@gmail.com>
 * Copyright (c) 2022.
 */

/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-10-22 09:09:37
 * @modify date 2022-10-22 09:09:37
 * @desc [description]
 */
package request

type CreateItemRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Qty         uint   `json:"qty"`
	CreatedBy   *uint  `json:"created_by"`
}

type GetItemPaginate struct {
	BaseRequest
}

type UpdateItemRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
	Qty         uint   `json:"qty"`
	Id          uint   `json:"id"`
}
