/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-07-03 20:58:28
 * @modify date 2022-07-03 20:58:28
 * @desc [description]
 */
package route

import (
	"samsul96maarif/github.com/go-api-app/handler"

	"github.com/gorilla/mux"
)

type ApiRoute struct {
	R       *mux.Router
	Handler *handler.Handler
}
