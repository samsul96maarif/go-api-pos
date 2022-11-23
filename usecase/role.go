/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-10-21 11:46:49
 * @modify date 2022-10-21 11:46:49
 * @desc [description]
 */

package usecase

import (
	"context"
	"samsul96maarif/github.com/go-api-app/model"
	"samsul96maarif/github.com/go-api-app/request"
)

func (usecase *Usecase) FindUserRole(ctx context.Context, req request.FindUserRoleRequest) (entity model.UserRole, err error) {
	entity, err = usecase.repo.FindUserRole(ctx, map[string]interface{}{"user_id": req.UserId, "role_id": req.RoleId}, "created_at")
	return entity, err
}
