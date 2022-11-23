/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-10-22 09:09:37
 * @modify date 2022-10-22 09:09:37
 * @desc [description]
 */
package request

type FindUserRoleRequest struct {
	UserId uint `json:"user_id" validate:"required"`
	RoleId uint `json:"role_id" validate:"required"`
}
