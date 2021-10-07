package users

import (
	"context"
	"kampus_merdeka/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) *MysqlUserRepository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Login(ctx context.Context, email string, password string) (users.Domain, error) {
	var user Users
	result := rep.Conn.First(&user, "email = ? AND password = ?",
		email, password)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil

}

func (rep *MysqlUserRepository) Register(ctx context.Context, user *users.Domain) (users.Domain, error) {
	// return user.ToDomain(), nil
	newUser := Users{
		Id:       user.Id,
		Email:    user.Email,
		Address:  user.Address,
		Name:     user.Name,
		Password: user.Password,
	}
	result := rep.Conn.Create(&newUser)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return newUser.ToDomain(), nil

}
