// utils/test_db.go
package utils

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupRouterAndMockDB(t *testing.T) (*gin.Engine, *gorm.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	dialector := postgres.New(postgres.Config{
		Conn:       sqldb,
		DriverName: "postgres",
	})

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatal(err)
	}

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	return router, db, mock
}