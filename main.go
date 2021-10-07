package main

import (
	"kampus_merdeka/app/routes"
	_sellerUsecase "kampus_merdeka/business/sellers"
	_userUsecase "kampus_merdeka/business/users"
	_sellerController "kampus_merdeka/controllers/sellers"
	_userController "kampus_merdeka/controllers/users"
	_sellerdb "kampus_merdeka/drivers/databases/sellers"
	_userdb "kampus_merdeka/drivers/databases/users"
	_mysqlDriver "kampus_merdeka/drivers/mysql"
	"time"

	_middleware "kampus_merdeka/app/middlewares"
	_sellerRepository "kampus_merdeka/drivers/databases/sellers"
	_userRepository "kampus_merdeka/drivers/databases/users"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_userdb.Users{})
	db.AutoMigrate(&_sellerdb.Sellers{})
}

func main() {
	// init koneksi databse
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	Conn := configDB.InitialDB()
	DbMigrate(Conn)

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userRepository.NewMysqlUserRepository(Conn)
	userUseCase := _userUsecase.NewUserUsecase(userRepository, timeoutContext)
	userController := _userController.NewUserController(userUseCase)

	sellerRepository := _sellerRepository.NewMysqlSellerRepository(Conn)
	sellerUseCase := _sellerUsecase.NewSellerUsecase(sellerRepository, timeoutContext)
	sellerController := _sellerController.NewSellerController(sellerUseCase)

	routesInit := routes.ControllerList{
		JwtConfig:        configJWT.Init(),
		UserController:   *userController,
		SellerController: *sellerController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}
