package database

import (
	"fmt"
	"service-management/internal/pkg/config"
	"service-management/internal/pkg/models/customers"
	"service-management/internal/pkg/models/services"
	"service-management/internal/pkg/models/users"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB  *gorm.DB
	err error
)

type Database struct {
	*gorm.DB
}

// SetupDB opens a database and saves the reference to `Database` struct.
func SetupDB() {
	var db = DB

	configuration := config.GetConfig()

	database := configuration.Database.Dbname
	username := configuration.Database.Username
	password := configuration.Database.Password
	host := configuration.Database.Host
	port := configuration.Database.Port

	db, err = gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" dbname="+database+"  sslmode=disable password="+password)
	if err != nil {
		fmt.Println("db err: ", err)
	}

	// Change this to true if you want to see SQL queries
	db.LogMode(false)
	db.DB().SetMaxIdleConns(configuration.Database.MaxIdleConns)
	db.DB().SetMaxOpenConns(configuration.Database.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(configuration.Database.MaxLifetime) * time.Second)
	DB = db
	migration()
}

func GetDB() *gorm.DB {
	return DB
}

// Auto migrate project models
func migration() {
	DB.AutoMigrate(&users.UserRole{})
	DB.AutoMigrate(&users.User{}).AddForeignKey("role_id", "user_roles(id)", "CASCADE", "CASCADE")
	DB.AutoMigrate(&customers.Customer{})
	DB.AutoMigrate(&customers.CustomerAddress{}).AddForeignKey("customer_id", "customers(id)", "CASCADE", "CASCADE")
	DB.AutoMigrate(&services.Service{})
}
