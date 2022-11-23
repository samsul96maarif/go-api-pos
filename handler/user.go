/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-10-21 13:58:56
 * @modify date 2022-10-21 13:58:56
 * @desc [description]
 */
package handler

import (
	"encoding/json"
	"net/http"
	"samsul96maarif/github.com/go-api-app/request"

	"github.com/go-playground/validator"
)

func (handler *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload request.CreateUser
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		writeError(w, err)
		return
	}
	err = validator.New().Struct(payload)
	if err != nil {
		writeError(w, err)
		return
	}
	res, err := handler.BE.Usecase.CreateUser(r.Context(), payload)
	if err != nil {
		writeError(w, err)
		return
	}
	writeSuccess(w, res, "Success", ResponseMeta{
		HttpStatus: http.StatusCreated,
	})
}
