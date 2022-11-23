package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"samsul96maarif/github.com/go-api-app/lib"
	"samsul96maarif/github.com/go-api-app/model"
)

func (usecase *Usecase) CreateSuperAdmin(ctx context.Context) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte("localhost"), bcrypt.MinCost)
	user := model.User{
		Email:    "super-admin@maarif.id",
		Password: string(hash),
	}
	erro := usecase.repo.Transaction(ctx, func(ctx context.Context) error {
		err := usecase.repo.CreateUser(ctx, &user)
		if err != nil {
			return err
		}
		userRole := model.UserRole{
			UserId: user.Id,
			RoleId: lib.ROLE_SUPER_ADMIN_ID,
		}
		err = usecase.repo.CreateUserRole(ctx, &userRole)
		return err
	})
	return erro
}
