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

	"golang.org/x/crypto/bcrypt"
)

func (usecase *Usecase) CreateUser(ctx context.Context, req request.CreateUser) (entity model.User, err error) {
	var hash []byte
	password := req.Password
	if password == "" {
		password = "password"
	}
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return entity, err
	}
	entity = model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}
	err = usecase.repo.CreateUser(ctx, &entity)
	return entity, err
}

func (usecase *Usecase) FindUser(ctx context.Context, req request.FindUserRequest) (entity model.User, err error) {
	if req.Id != 0 {
		entity, err = usecase.repo.FindUser(ctx, map[string]interface{}{"id": req.Id}, "created_at")
	}
	if req.Email != "" {
		entity, err = usecase.repo.FindUser(ctx, map[string]interface{}{"email": req.Email}, "created_at")
	}
	return entity, err
}
