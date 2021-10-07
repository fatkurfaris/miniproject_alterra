package requests

import "kampus_merdeka/business/sellers"

type SellerLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *SellerLogin) ToDomain() sellers.Domain {
	return sellers.Domain{
		Email:    user.Email,
		Password: user.Password,
	}
}
