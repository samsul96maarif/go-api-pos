/**
 * @author [Samsul Ma'arif]
 * @email [samsulma828@gmail.com]
 * @create date 2022-07-03 13:56:24
 * @modify date 2022-07-03 13:56:24
 * @desc [description]
 */
/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2022.
 */

package repository

import (
	"context"

	"samsul96maarif/github.com/go-api-app/lib"
	"samsul96maarif/github.com/go-api-app/model"
	// "github.com/golang/protobuf/ptypes/any"
)

type Repository struct {
	db *lib.Database
}

func NewRepository(db *lib.Database) Repository { return Repository{db: db} }

func (repo *Repository) Transaction(ctx context.Context, fn func(context.Context) error) (err error) {
	trx := repo.db.Begin()
	ctx = context.WithValue(ctx, "Trx", &lib.Database{DB: trx})
	if err = fn(ctx); err != nil {
		trx.Rollback()
	} else {
		err = trx.Commit().Error
	}
	return err
}
func cok[V int | string](s V) {

}

type b interface {
	model.User | model.Message
}
