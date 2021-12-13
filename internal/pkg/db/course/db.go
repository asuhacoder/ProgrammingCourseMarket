package course

import (
	"time"

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
		DSN: "host=course_db user=gorm password=gorm dbname=gorm port=5433 sslmode=disable TimeZone=Asia/Tokyo",
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
	DB.AutoMigrate(&Course{})
}

// User Model
type Course struct {
	UUID       uuid.UUID `gorm:"primaryKey; unique; type:uuid;"`
	USER_ID    uuid.UUID `gorm:"not null"`
	TITLE      string    `gorm:"not null"`
	INTRODUCE  string    `gorm:"not null"`
	IMAGE      string
	PRICE      int       `gorm:"not null"`
	PUBLISHED  bool      `gorm:"not null"`
	CREATED_AT time.Time `gorm:"autoCreateTime"`
}
