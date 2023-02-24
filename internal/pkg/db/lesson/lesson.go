package lesson

import (
	"log"

	uuid "github.com/gofrs/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	LessonDB *gorm.DB
	err      error
)

func LessonInit() {
	LessonDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "host=lesson_db user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Tokyo",
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	lessonAutoMigration()
}

func LessonClose() {
	dbSQL, err := LessonDB.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}

func lessonAutoMigration() {
	err := LessonDB.AutoMigrate(&Lesson{})
	if err != nil {
		log.Printf("failed to AutoMigrate: %v", err)
	}
}

type Lesson struct {
	UUID            uuid.UUID `gorm:"primaryKey; unique; type:uuid;"`
	USER_ID         uuid.UUID `gorm:"not null"`
	COURSE_ID       uuid.UUID `gorm:"not null"`
	SEQUENCE_NUMBER int64     `gorm:"not null"`
	TITLE           string    `gorm:"not null"`
	INTRODUCTION    string    `gorm:"not null"`
	BODY            string    `gorm:"not null"`
	DEFAULT_CODE    string    `gorm:"not null"`
	ANSWER_CODE     string    `gorm:"not null"`
	LANGUAGE        string    `gorm:"not null"`
}
