/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-07-06 21:27:13
 * @modify date 2022-07-06 21:27:13
 * @desc [description]
 */

package handler

import (
	"encoding/json"
	"net/http"
	"samsul96maarif/github.com/go-api-app/lib"
	"samsul96maarif/github.com/go-api-app/request"

	"github.com/go-playground/validator"
)

func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var payload request.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		writeError(w, lib.ErrorInvalidParameter)
		return
	}
	err = validator.New().Struct(payload)
	if err != nil {
		writeError(w, err)
		return
	}
	res, err := handler.BE.Usecase.Register(r.Context(), payload)
	if err != nil {
		writeError(w, err)
		return
	}
	writeSuccess(w, res, "Success", ResponseMeta{
		HttpStatus: http.StatusCreated,
	})
}

func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var payload request.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		writeError(w, err)
		return
	}
	//err = validator.New().Struct(payload)
	//if err != nil {
	//	writeError(w, err)
	//	return
	//}
	//fmt.Printf("%+v \n", payload)
	res, err := handler.BE.Usecase.Login(r.Context(), payload)
	if err != nil {
		writeError(w, err)
		return
	}
	writeSuccess(w, res, "Success", ResponseMeta{HttpStatus: http.StatusOK})
	return
}
