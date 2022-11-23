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
	"samsul96maarif/github.com/go-api-app/request"

	"samsul96maarif/github.com/go-api-app/lib"
	"samsul96maarif/github.com/go-api-app/lib/logger"
	"samsul96maarif/github.com/go-api-app/model"

	"gorm.io/gorm"
)

func (repo *Repository) FindItem(ctx context.Context, where map[string]interface{}, order_by string) (entity model.Item, err error) {
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

func (repo *Repository) CreateItem(ctx context.Context, entity *model.Item) (err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}
	err = tx.Create(entity).Error
	if err != nil {
		logger.Error(ctx, "Error create item", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "create", "user"},
		})
	}
	return err
}

func (repo *Repository) GetItemPaginate(ctx context.Context, where map[string]interface{}, req request.BaseRequest, order_by string) (entities []model.Item, err error) {
	page := req.GetPage()
	limit := req.GetLimit()
	offset := int((page - 1) * limit)
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.Where(where).Order(order_by).Offset(offset).Limit(int(limit)).Find(&entities).Error
	if err != nil {
		logger.Error(ctx, "Error GetItemPaginate", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "db", "item", "get"},
		})
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return entities, err
		}
	}
	return entities, nil
}

func (repo *Repository) GetItemCount(ctx context.Context, where map[string]interface{}) (total int64, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	statement := tx.Model(&model.Item{}).Where(where)

	err = statement.Count(&total).Error
	if err != nil {
		go logger.Error(ctx, "error count item", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres"},
		})
	}

	return total, err
}
