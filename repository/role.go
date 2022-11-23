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

func (repo *Repository) GetRole(ctx context.Context, where map[string]interface{}, order_by string) (entities []model.Role, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.Where(where).Find(&entities).Order(order_by).Error
	if err != nil {
		logger.Error(ctx, "Error get roles", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "get", "role"},
		})
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return entities, err
		}
	}
	return entities, nil
}

func (repo *Repository) GetUserRole(ctx context.Context, where map[string]interface{}, order_by string) (entities []model.UserRole, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.Where(where).Find(&entities).Order(order_by).Error
	if err != nil {
		logger.Error(ctx, "Error get user_roles", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "get", "role"},
		})
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return entities, err
		}
	}
	return entities, nil
}

func (repo *Repository) FindUserRole(ctx context.Context, where map[string]interface{}, order_by string) (entity model.UserRole, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}
	err = tx.Where(where).Limit(1).Find(&entity).Order(order_by).Error
	if err != nil {
		logger.Error(ctx, "Error get userRole", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "find", "userRole"},
		})
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return entity, err
		}
	}
	return entity, nil
}

func (repo *Repository) CreateUserRole(ctx context.Context, entity *model.UserRole) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}
	err := tx.Create(entity).Error
	if err != nil {
		logger.Error(ctx, "Error CreateUserRole", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "user_roles", "create"},
		})
	}
	return err
}
