package lesson

import (
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
		DSN: "host=lesson_db user=gorm password=gorm dbname=gorm port=5434 sslmode=disable TimeZone=Asia/Tokyo",
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
	DB.AutoMigrate(&Lesson{})
}

// Lesson Model
type Lesson struct {
	UUID         uuid.UUID  `gorm:"primaryKey; unique; type:uuid;"`
	USER_ID      uuid.UUID  `gorm:"not null"`
	COURSE_ID    uuid.UUID  `gorm:"not null"`
	TITLE        string     `gorm:"not null"`
	INTRODUCTION string     `gorm:"not null"`
	BODY         string     `gorm:"not null"`
	DEFAULT_CODE string     `gorm:"not null"`
	ANSWER_CODE  string     `gorm:"not null"`
	TestCase     []TestCase `gorm:"not null"`
	LANGUAGE     string     `gorm:"not null"`
}

// Case Model
type TestCase struct {
	LESSON_ID uuid.UUID
	INPUT     string
	OUTPUT    string
}
