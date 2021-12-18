package user

import (
	"log"

	uuid "github.com/gofrs/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DB instance
	DB  *gorm.DB
	err error
)

// Init DB
func Init() {
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "host=user_db user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Tokyo",
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	autoMigration()
}

// Close DB
func Close() {
	dbSQL, err := DB.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}

func autoMigration() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		log.Printf("failed to AutoMigrate: %v", err)
	}
}

// User Model
type User struct {
	UUID       uuid.UUID `gorm:"primaryKey; unique; type:uuid;"`
	EMAIL      string    `gorm:"unique"`
	PERMISSION string    `gorm:"not null"`
	PASSWORD   string    `gorm:"not null"`
}
