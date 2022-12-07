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
	"samsul96maarif/github.com/go-api-app/lib"
	"samsul96maarif/github.com/go-api-app/lib/logger"
	"samsul96maarif/github.com/go-api-app/model"
	"samsul96maarif/github.com/go-api-app/request"
	"samsul96maarif/github.com/go-api-app/response"
)

func (usecase *Usecase) CreateItemPrice(ctx context.Context, req request.CreateItemPriceRequest) (entity model.ItemPrice, err error) {
	err = usecase.repo.Transaction(ctx, func(ctx context.Context) (erro error) {
		existingEntity, _ := usecase.repo.FindItemPrice(ctx, map[string]interface{}{
			"item_id":    req.ItemId,
			"is_default": true,
		}, "created_at asc")
		isDefault := req.IsDefault

		if isDefault {
			if existingEntity.Id != 0 {
				erro = usecase.repo.UpdateItemPrice(ctx, &existingEntity, map[string]interface{}{"is_default": false})
				if erro != nil {
					logger.Error(ctx, "Error UpdateItemPrice", map[string]interface{}{
						"error": err,
						"tags":  []string{"update", "item_prices"},
					})
					return erro
				}
			}
		}
		if !isDefault && existingEntity.Id == 0 {
			isDefault = true
		}

		currentUser := ctx.Value("User").(lib.MyClaim)
		code := request.GenerateCode(req.Name)
		entity = model.ItemPrice{
			Name:      req.Name,
			Price:     req.Price,
			Code:      code,
			ItemId:    req.ItemId,
			CreatedBy: &currentUser.UserId,
			IsDefault: isDefault,
		}
		erro = usecase.repo.CreateItemPrice(ctx, &entity)
		return erro
	})

	return entity, err
}

func (usecase *Usecase) FindItemPrice(ctx context.Context, id uint) (entity model.ItemPrice, err error) {
	entity, err = usecase.repo.FindItemPrice(ctx, map[string]interface{}{"id": id}, "created_at asc")
	return entity, err
}

func (usecase *Usecase) DeleteItemPrice(ctx context.Context, id uint) error {
	defaultPrice, _ := usecase.repo.FindItemPrice(ctx, map[string]interface{}{"id": id, "is_default": true}, "created_at asc")
	if defaultPrice.Id != 0 {
		return lib.ErrorCannotDeleteDefaultPrice
	}
	err := usecase.repo.DeleteItemPrice(ctx, map[string]interface{}{"id": id})
	return err
}

func (u *Usecase) UpdateItemPrice(ctx context.Context, req request.UpdateItemPriceRequest) (entity model.ItemPrice, err error) {
	entity, err = u.repo.FindItemPrice(ctx, map[string]interface{}{"id": req.Id}, "created_at asc")
	if err != nil {
		return entity, err
	}
	if entity.Id == 0 {
		return entity, lib.ErrorNotFound
	}
	var defaultPrice model.ItemPrice
	if req.IsDefault && !entity.IsDefault {
		defaultPrice, _ = u.repo.FindItemPrice(ctx, map[string]interface{}{
			"item_id":    entity.ItemId,
			"is_default": true,
		}, "created_at asc")
	}
	err = u.repo.Transaction(ctx, func(ctx context.Context) (erro error) {
		if defaultPrice.Id != 0 && req.IsDefault {
			erro = u.repo.UpdateItemPrice(ctx, &defaultPrice, map[string]interface{}{"is_default": false})
			if erro != nil {
				return erro
			}

		}
		erro = u.repo.UpdateItemPrice(ctx, &entity, map[string]interface{}{
			"name":       req.Name,
			"price":      req.Price,
			"code":       req.Code,
			"is_default": req.IsDefault,
		})
		return erro
	})
	if err != nil {
		logger.Error(ctx, "Error UpdateItemPrice", map[string]interface{}{
			"error": err,
			"tags":  []string{"usecase", "update", "item_prices"},
		})
	}
	return entity, err
}

func (u *Usecase) GetItemPricePaginate(ctx context.Context, req request.GetItemPricePaginate) (res response.GetItemPricePaginateResponse, err error) {
	var entities []model.ItemPrice
	var total int64
	entities, err = u.repo.GetItemPricePaginate(ctx, map[string]interface{}{}, req.BaseRequest, "created_at asc")
	if err != nil {
		logger.Error(ctx, "Error GetItemPricePaginate", map[string]interface{}{
			"error": err,
			"tags":  []string{"item_prices", "repo", "get"},
		})
		return res, err
	}
	total, err = u.repo.GetItemPriceCount(ctx, map[string]interface{}{}, req.Keyword)
	if err != nil {
		logger.Error(ctx, "Error GetItemPriceCount", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "item_prices", "get_count"},
		})
		return res, err
	}
	res.ItemPrices = entities
	res.Total = total
	res.Limit = req.Limit
	res.Page = req.Page
	return res, err
}
