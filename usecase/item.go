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
	"fmt"
	"samsul96maarif/github.com/go-api-app/lib"
	"samsul96maarif/github.com/go-api-app/lib/logger"
	"samsul96maarif/github.com/go-api-app/model"
	"samsul96maarif/github.com/go-api-app/request"
	"samsul96maarif/github.com/go-api-app/response"
)

func (usecase *Usecase) CreateItem(ctx context.Context, req request.CreateItemRequest) (entity model.Item, err error) {
	currentUser := ctx.Value("User").(lib.MyClaim)
	fmt.Printf("current %+v \n", currentUser)
	entity = model.Item{
		Name:        req.Name,
		Description: req.Description,
		Qty:         req.Qty,
		CreatedBy:   &currentUser.UserId,
	}
	err = usecase.repo.CreateItem(ctx, &entity)
	return entity, err
}

func (usecase *Usecase) GetItemPaginate(ctx context.Context, req request.GetItemPaginate) (res response.GetItemPaginateResponse, err error) {
	where := make(map[string]interface{})
	if req.Keyword != "" {
		search := fmt.Sprintf("%%%s%%", req.Keyword)
		where["name LIKE ?"] = search
	}
	var entities []model.ItemWithDefaultPrice
	var total int64
	entities, err = usecase.repo.GetItemPaginate(ctx, where, req.BaseRequest, "name asc")
	if err != nil {
		logger.Error(ctx, "Error GetItemPaginate", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "usecase", "items", "get"},
		})
	}
	total, err = usecase.repo.GetItemCount(ctx, where, req.Keyword)
	if err != nil {
		logger.Error(ctx, "Error GetItemCount", map[string]interface{}{
			"error": err,
			"tags":  []string{"repo", "items", "usecase", "count"},
		})
	}

	res.Items = entities
	res.Total = total
	res.Limit = req.Limit
	res.Page = req.Page
	return res, err
}

func (usecase *Usecase) FindItem(ctx context.Context, id uint) (entity model.Item, err error) {
	entity, err = usecase.repo.FindItem(ctx, map[string]interface{}{"id": id}, "created_at asc")
	return entity, err
}

func (u *Usecase) DeleteItem(ctx context.Context, id uint) error {
	err := u.repo.DeleteItem(ctx, map[string]interface{}{"id": id})
	return err
}

func (usecase *Usecase) UpdateItem(ctx context.Context, req request.UpdateItemRequest) (entity model.Item, err error) {
	entity, err = usecase.repo.UpdateItem(ctx, map[string]interface{}{"id": req.Id}, map[string]interface{}{
		"name":        req.Name,
		"description": req.Description,
		"qty":         req.Qty,
	})
	return entity, err
}
