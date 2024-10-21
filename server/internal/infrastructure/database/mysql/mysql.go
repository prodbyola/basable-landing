package mysql

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"basable/internal/infrastructure/database/entities"
	"basable/pkg/utils/logger"
)

func ConnectToDB(host string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(host), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		log.Fatal("::failed to connect to mysql database with error", err)
		return nil
	}
	logger.Slog().Info("::successfully connected to mysql database")
	return db
}

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&entities.Collaborator{}); err != nil {
		log.Fatal("::failed to migrate database with error", err)
		return
	}
	logger.Slog().Info("::successfully migrated database")
}
