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

import (
	"strconv"
	"strings"
	"time"
)

type CreateItemPriceRequest struct {
	Name      string `json:"name" validate:"required"`
	Code      string `json:"code"`
	Price     uint   `json:"price"`
	ItemId    uint   `json:"item_id"`
	CreatedBy *uint  `json:"created_by"`
	IsDefault bool   `json:"is_default"`
}

type UpdateItemPriceRequest struct {
	Name      string `json:"name,omitempty"`
	Code      string `json:"code"`
	Price     uint   `json:"price"`
	Id        uint   `json:"id"`
	IsDefault bool   `json:"is_default"`
}

type GetItemPricePaginate struct {
	BaseRequest
}

func GenerateCode(name string) (code string) {
	name = strings.ToLower(name)
	slugName := strings.ReplaceAll(name, " ", "-")
	timeUnix := time.Now().Unix()
	timeStr := strconv.Itoa(int(timeUnix))
	return slugName + timeStr
}
