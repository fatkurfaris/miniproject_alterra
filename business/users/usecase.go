package users

import (
	"context"
	"errors"
	"time"
)

type UserUsecase struct {
	// ConfigJWT      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repo Repository, timeout time.Duration) Usecase {
	return &UserUsecase{
		// ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *UserUsecase) Register(ctx context.Context, domain Domain) (Domain, error) {

	if domain.Name == "" {
		return Domain{}, errors.New("name empty")
	}

	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}

	if domain.Address == "" {
		return Domain{}, errors.New("address empty")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}

	user, err := uc.Repo.Register(ctx, &domain)

	if err != nil {
		return Domain{}, err
	}

	if err != nil {
		return Domain{}, err
	}

	if err != nil {
		return Domain{}, err
	}

	if err != nil {
		return Domain{}, err
	}

	return user, nil

}

// core bisinis login
func (uc *UserUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {

	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	// var err error
	// domain.Password, err = encrypt.Hash(domain.Password)

	// if err != nil {
	// 	return Domain{}, err
	// }

	user, err := uc.Repo.Login(ctx, domain.Email, domain.Password)

	if err != nil {
		return Domain{}, err
	}

	// user.Token, err = uc.ConfigJWT.GenerateToken(user.Id)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
