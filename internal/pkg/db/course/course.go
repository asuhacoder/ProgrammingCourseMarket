package course

import (
	"log"
	"time"

	uuid "github.com/gofrs/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "host=course_db user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Tokyo",
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	autoMigration()
}

func Close() {
	dbSQL, err := DB.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}

func autoMigration() {
	err := DB.AutoMigrate(&Course{})
	if err != nil {
		log.Printf("failed to AutoMigrate: %v", err)
	}
}

type Course struct {
	UUID         uuid.UUID `gorm:"primaryKey; unique; type:uuid;"`
	USER_ID      uuid.UUID `gorm:"not null"`
	TITLE        string    `gorm:"not null"`
	INTRODUCTION string    `gorm:"not null"`
	IMAGE        string
	PRICE        int       `gorm:"not null"`
	IS_PUBLIC    bool      `gorm:"not null"`
	CREATED_AT   time.Time `gorm:"autoCreateTime"`
}
