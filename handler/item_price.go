/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2022.
 */

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"samsul96maarif/github.com/go-api-app/lib"
	"samsul96maarif/github.com/go-api-app/request"
	"strconv"
)

func (handler *Handler) CreateItemPrice(w http.ResponseWriter, r *http.Request) {
	var payload request.CreateItemPriceRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		writeError(w, err)
		return
	}
	currentUser := r.Context().Value("User").(lib.MyClaim)
	fmt.Printf("handler %+v \n", currentUser)
	entity, err := handler.BE.Usecase.CreateItemPrice(r.Context(), payload)
	if err != nil {
		writeError(w, err)
		return
	}
	writeSuccess(w, entity, "Succeed", ResponseMeta{HttpStatus: http.StatusCreated})
	return
}

func (handler *Handler) FindItemPrice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)
	entity, err := handler.BE.Usecase.FindItemPrice(r.Context(), uint(id))
	if err != nil {
		writeError(w, err)
		return
	}
	writeSuccess(w, entity, "Succeed", ResponseMeta{
		HttpStatus: http.StatusOK,
	})
	return
}

func (handler *Handler) GetItemPricePaginate(w http.ResponseWriter, r *http.Request) {
	payload := request.GetItemPricePaginate{
		BaseRequest: GenerateBaseRequest(r),
	}
	res, err := handler.BE.Usecase.GetItemPricePaginate(r.Context(), payload)
	if err != nil {
		writeError(w, err)
		return
	}
	writeSuccess(w, res, "Succeed", ResponseMeta{
		HttpStatus: http.StatusOK,
	})
}

func (handler *Handler) UpdateItemPrice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		writeError(w, lib.InvalidParameterError("id", "id invalid"+err.Error()))
		return
	}
	var payload request.UpdateItemPriceRequest
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		writeError(w, err)
		return
	}
	payload.Id = uint(idInt)
	entity, err := handler.BE.Usecase.UpdateItemPrice(r.Context(), payload)
	if err != nil {
		writeError(w, err)
		return
	}
	writeSuccess(w, entity, "Succeed", ResponseMeta{
		HttpStatus: http.StatusOK,
	})
}

func (handler *Handler) DeleteItemPrice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		writeError(w, lib.InvalidParameterError("id", "id invalid "+err.Error()))
		return
	}
	err = handler.BE.Usecase.DeleteItemPrice(r.Context(), uint(idInt))
	if err != nil {
		writeError(w, err)
		return
	}
	writeSuccess(w, "", "Succeed", ResponseMeta{
		HttpStatus: http.StatusNoContent,
	})
}
