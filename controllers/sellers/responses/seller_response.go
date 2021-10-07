package responses

import (
	"kampus_merdeka/business/sellers"
	"time"
)

type SellerResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain sellers.Domain) SellerResponse {
	return SellerResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Address:   domain.Address,
		Token:     domain.Token,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
