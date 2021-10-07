package sellers

import (
	"context"
	"kampus_merdeka/business/sellers"

	"gorm.io/gorm"
)

type MysqlSellerRepository struct {
	Conn *gorm.DB
}

func NewMysqlSellerRepository(conn *gorm.DB) *MysqlSellerRepository {
	return &MysqlSellerRepository{
		Conn: conn,
	}
}

func (rep *MysqlSellerRepository) Login(ctx context.Context, email string, password string) (sellers.Domain, error) {
	var seller Sellers
	result := rep.Conn.First(&seller, "email = ? AND password = ?",
		email, password)

	if result.Error != nil {
		return sellers.Domain{}, result.Error
	}

	return seller.ToDomain(), nil

}

func (rep *MysqlSellerRepository) Register(ctx context.Context, user *sellers.Domain) (sellers.Domain, error) {
	// return user.ToDomain(), nil
	newSeller := Sellers{
		Id:       user.Id,
		Email:    user.Email,
		Address:  user.Address,
		Name:     user.Name,
		Password: user.Password,
	}
	result := rep.Conn.Create(&newSeller)
	if result.Error != nil {
		return sellers.Domain{}, result.Error
	}

	return newSeller.ToDomain(), nil

}
