package sellers

import (
	"context"
	"errors"
	"time"
)

type SellerUsecase struct {
	// ConfigJWT      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewSellerUsecase(repo Repository, timeout time.Duration) Usecase {
	return &SellerUsecase{
		// ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *SellerUsecase) Register(ctx context.Context, domain Domain) (Domain, error) {

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

	seller, err := uc.Repo.Register(ctx, &domain)

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

	return seller, nil

}

// core bisinis login
func (uc *SellerUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {

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

	seller, err := uc.Repo.Login(ctx, domain.Email, domain.Password)

	if err != nil {
		return Domain{}, err
	}

	// user.Token, err = uc.ConfigJWT.GenerateToken(user.Id)

	if err != nil {
		return Domain{}, err
	}

	return seller, nil
}
