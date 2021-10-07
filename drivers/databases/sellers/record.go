package sellers

import (
	"kampus_merdeka/business/sellers"
	"time"

	"gorm.io/gorm"
)

type Sellers struct {
	Id        int    `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Name      string
	Address   string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (seller *Sellers) ToDomain() sellers.Domain {
	return sellers.Domain{
		Id:        seller.Id,
		Name:      seller.Name,
		Email:     seller.Email,
		Address:   seller.Address,
		Password:  seller.Password,
		CreatedAt: seller.CreatedAt,
		UpdatedAt: seller.UpdatedAt,
	}
}

func ToListDomain(data []Sellers) (result []sellers.Domain) {
	result = []sellers.Domain{}
	for _, seller := range data {
		result = append(result, seller.ToDomain())
	}
	return
}

func FromDomain(domain sellers.Domain) Sellers {
	return Sellers{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Address:   domain.Address,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
