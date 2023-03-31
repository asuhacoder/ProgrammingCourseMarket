package test_case

import (
	"log"
	"os"

	uuid "github.com/gofrs/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	CaseDB *gorm.DB
	err    error
)

func CaseInit() {
	host := os.Getenv("TEST_CASE_DB_HOST")
	CaseDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "host=" + host + " user=gorm password=gormpassword dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Tokyo",
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	caseAutoMigration()
}

func CaseClose() {
	dbSQL, err := CaseDB.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}

func caseAutoMigration() {
	err := CaseDB.AutoMigrate(&TestCase{})
	if err != nil {
		log.Printf("failed to AutoMigrate: %v", err)
	}
}

type TestCase struct {
	UUID      uuid.UUID `gorm:"primaryKey; unique; type:uuid;"`
	USER_ID   uuid.UUID `gorm:"not null"`
	LESSON_ID uuid.UUID `gorm:"not null"`
	INPUT     string    `gorm:"not null"`
	OUTPUT    string    `gorm:"not null"`
}
