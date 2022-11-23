/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-07-03 21:04:56
 * @modify date 2022-07-03 21:04:56
 * @desc [description]
 */
package route

import (
	"net/http"
	"samsul96maarif/github.com/go-api-app/lib"
)

func (route ApiRoute) AddItemRoute() {
	route.R.HandleFunc("/items", route.Handler.ProtectedByRolesMiddleware(http.HandlerFunc(route.Handler.CreateItem), []int{lib.ROLE_SUPER_ADMIN_ID}).ServeHTTP).Methods("POST")
	route.R.HandleFunc("/items", route.Handler.PublicMiddleware(http.HandlerFunc(route.Handler.GetItemPaginate)).ServeHTTP).Methods("GET")
}
