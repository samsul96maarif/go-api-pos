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

func (route ApiRoute) AddItemPriceRoute() {
	route.R.HandleFunc("/item-prices", route.Handler.ProtectedByRolesMiddleware(http.HandlerFunc(route.Handler.CreateItemPrice), []int{lib.ROLE_SUPER_ADMIN_ID}).ServeHTTP).Methods("POST")
	route.R.HandleFunc("/item-prices", route.Handler.ProtectedByRolesMiddleware(http.HandlerFunc(route.Handler.GetItemPricePaginate), []int{lib.ROLE_SUPER_ADMIN_ID}).ServeHTTP).Methods("GET")
	route.R.HandleFunc("/item-prices/{id:[0:9]+}", route.Handler.ProtectedByRolesMiddleware(http.HandlerFunc(route.Handler.UpdateItemPrice), []int{lib.ROLE_SUPER_ADMIN_ID}).ServeHTTP).Methods("PUT")
	route.R.HandleFunc("/item-prices/{id:[0:9]+}", route.Handler.ProtectedByRolesMiddleware(http.HandlerFunc(route.Handler.DeleteItemPrice), []int{lib.ROLE_SUPER_ADMIN_ID}).ServeHTTP).Methods("DELETE")
	route.R.HandleFunc("/item-prices/{id:[0:9]+}", route.Handler.ProtectedByRolesMiddleware(http.HandlerFunc(route.Handler.FindItemPrice), []int{lib.ROLE_SUPER_ADMIN_ID, lib.ROLE_ADMIN_ID}).ServeHTTP).Methods("GET")
}
