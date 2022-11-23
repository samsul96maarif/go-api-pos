/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-07-03 20:52:18
 * @modify date 2022-07-03 20:52:18
 * @desc [description]
 */
package handler

import (
	"encoding/json"
	"net/http"
	goapiapp "samsul96maarif/github.com/go-api-app"
	"samsul96maarif/github.com/go-api-app/lib"
)

type Handler struct{ BE *goapiapp.BE }

type ResponseMeta struct {
	HttpStatus int   `json:"http_status"`
	Total      *uint `json:"total,omitempty"`
	PerPage    *uint `json:"per_page,omitempty"`
	Page       *uint `json:"page,omitempty"`
}

type ResponseBody struct {
	Meta    ResponseMeta `json:"meta"`
	Data    interface{}  `json:"data,omitempty"`
	Message string       `json:"message,omitempty"`
}

type ErrorInfo struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
	Code    int    `json:"code"`
}

type ErrorBody struct {
	Errors []ErrorInfo `json:"errors"`
	Meta   interface{} `json:"meta"`
}

func NewHandler(be *goapiapp.BE) Handler {
	return Handler{BE: be}
}

func writeError(w http.ResponseWriter, err error) {
	var resp interface{}
	code := http.StatusInternalServerError

	switch errOrig := err.(type) {
	case lib.CustomError:
		resp = ErrorBody{
			Errors: []ErrorInfo{
				{
					Message: errOrig.Message,
					Code:    errOrig.Code,
					Field:   errOrig.Field,
				},
			},
			Meta: ResponseMeta{
				HttpStatus: errOrig.HttpCode,
			},
		}

		code = int(errOrig.HttpCode)
	default:
		resp = ResponseBody{
			Message: "Internal Server Error",
			Meta: ResponseMeta{
				HttpStatus: code,
			},
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

func writeSuccess(w http.ResponseWriter, data interface{}, message string, meta ResponseMeta) {
	resp := ResponseBody{
		Message: message,
		Data:    data,
		Meta:    meta,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(meta.HttpStatus)
	json.NewEncoder(w).Encode(resp)
}

func writeResponse(w http.ResponseWriter, resp interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
