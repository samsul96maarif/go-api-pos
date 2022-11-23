/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-07-03 15:33:30
 * @modify date 2022-07-03 15:33:30
 * @desc [description]
 */
package repository

import (
	"context"
	"errors"

	"samsul96maarif/github.com/go-api-app/lib"
	"samsul96maarif/github.com/go-api-app/lib/logger"
	"samsul96maarif/github.com/go-api-app/model"

	"gorm.io/gorm"
)

func (repo *Repository) FindUser(ctx context.Context, where map[string]interface{}, order_by string) (entity model.User, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.Where(where).Limit(1).Find(&entity).Order(order_by).Error
	if err != nil {
		logger.Error(ctx, "Error find user", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "find", "user"},
		})
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return entity, err
		}
	}
	return entity, nil
}

func (repo *Repository) CreateUser(ctx context.Context, entity *model.User) (err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}
	err = tx.Create(entity).Error
	if err != nil {
		logger.Error(ctx, "Error create user", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "create", "user"},
		})
	}
	return err
}
