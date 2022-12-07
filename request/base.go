/*
 * Author: Samsul Ma'arfi <samsulma828@gmail.com>
 * Copyright (c) 2022.
 */

package request

type BaseRequest struct {
	Sorts   []string `json:"sorts,omitempty;query:sorts"`
	Keyword string   `json:"keyword,omitempty;query:keyword"`
	Page    uint     `json:"page,omitempty;query:page"`
	Limit   uint     `json:"limit,omitempty;query:limit"`
}

func (req *BaseRequest) GetPage() uint {
	if req.Page > 0 {
		return req.Page
	}
	return 1
}

func (req *BaseRequest) GetLimit() uint {
	if req.Limit > 0 {
		return req.Limit
	}
	return 100
}
