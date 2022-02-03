package infra

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DSN struct {
	Host string
	User string
	Port string
	DB   string
}

func NewDB() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "postgres",
		DSN:        "host=localhost user=postgres port=5432 dbname=test_mock_db sslmode=disable",
	}))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("failed to create mock db")
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}))
	if err != nil {
		log.Fatal("failed to open gorm")
	}
	return gormDB, mock
}
