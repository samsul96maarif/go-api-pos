/*
 * Author: Samsul Ma'arfi <samsulma828@gmail.com>
 * Copyright (c) 2022.
 */

package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"samsul96maarif/github.com/go-api-app/request"
	"strconv"
)

func (handler *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var payload request.CreateItemRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		writeError(w, err)
		return
	}
	entity, err := handler.BE.Usecase.CreateItem(r.Context(), payload)
	if err != nil {
		writeError(w, err)
		return
	}
	writeSuccess(w, entity, "Succeed", ResponseMeta{HttpStatus: http.StatusCreated})
	return
}

func (handler *Handler) FindItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)
	entity, err := handler.BE.Usecase.FindItem(r.Context(), uint(id))
	if err != nil {
		writeError(w, err)
		return
	}
	writeSuccess(w, entity, "Succeed", ResponseMeta{
		HttpStatus: http.StatusOK,
	})
	return
}

func (handler *Handler) GetItemPaginate(w http.ResponseWriter, r *http.Request) {
	var payload request.GetItemPaginate
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		writeError(w, err)
		return
	}
	entities, err := handler.BE.Usecase.GetItemPaginate(r.Context(), payload)
	if err != nil {
		writeError(w, err)
		return
	}
	writeSuccess(w, entities, "Succeed", ResponseMeta{
		HttpStatus: http.StatusOK,
	})
	return
}
