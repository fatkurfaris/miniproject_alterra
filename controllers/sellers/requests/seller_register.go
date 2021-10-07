package requests

import "kampus_merdeka/business/sellers"

type SellerRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

func (seller *SellerRegister) ToDomain() sellers.Domain {
	return sellers.Domain{
		Name:     seller.Name,
		Email:    seller.Email,
		Address:  seller.Address,
		Password: seller.Password,
	}
}
