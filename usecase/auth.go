package usecase

import (
	"context"
	"fmt"
	"os"
	"samsul96maarif/github.com/go-api-app/config"
	"samsul96maarif/github.com/go-api-app/lib"
	"samsul96maarif/github.com/go-api-app/model"
	"samsul96maarif/github.com/go-api-app/request"
	"samsul96maarif/github.com/go-api-app/response"

	"github.com/golang-jwt/jwt"

	"golang.org/x/crypto/bcrypt"
)

func (usecase *Usecase) Register(ctx context.Context, req request.RegisterRequest) (res response.UserResponse, err error) {
	if req.Password != req.PasswordConfirmation {
		return res, lib.InvalidParameterError("password_confirmation", "Invalid password confirmation")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return res, err
	}
	entity := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}
	err = usecase.repo.CreateUser(ctx, &entity)
	if err != nil {
		return res, err
	}
	res = response.ConvertUserModelToUserResponse(entity)
	return res, err
}

func (usecase *Usecase) Login(ctx context.Context, req request.LoginRequest) (res response.LoginResponse, err error) {
	entity, err := usecase.repo.FindUser(ctx, map[string]interface{}{"email": req.Email}, "created_at")
	if err != nil {
		return res, err
	}
	roles, err := usecase.repo.GetUserRole(ctx, map[string]interface{}{"user_id": entity.Id}, "created_at")
	if err != nil {
		return res, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(entity.Password), []byte(req.Password))
	if err != nil {
		return res, err
	}

	var userRoles []int
	for _, val := range roles {
		userRoles = append(userRoles, int(val.RoleId))
	}
	claims := lib.MyClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer: os.Getenv("APP_NAME"),
		},
		UserId: entity.Id,
		Email:  entity.Email,
		Roles:  userRoles,
	}
	fmt.Printf("%+v \n", claims)

	token := jwt.NewWithClaims(config.JWT_SIGNING_METHOD, claims)

	signedToken, err := token.SignedString(config.GetSignatureKey())
	if err != nil {
		return res, err
	}
	res.Token = signedToken
	return res, err
}
