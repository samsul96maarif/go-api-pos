/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-07-03 21:04:56
 * @modify date 2022-07-03 21:04:56
 * @desc [description]
 */
package route

import "net/http"

func (route ApiRoute) AddAuthRoute() {
	route.R.HandleFunc("/register", route.Handler.PublicMiddleware(http.HandlerFunc(route.Handler.Register)).ServeHTTP).Methods("POST")
	route.R.HandleFunc("/login", route.Handler.PublicMiddleware(http.HandlerFunc(route.Handler.Login)).ServeHTTP).Methods("POST")
}
