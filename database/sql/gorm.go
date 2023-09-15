package sql

import (
	"time"

	"github.com/Ho-J/base/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func NewGorm(conn string) *gorm.DB {

	logger := logger.New(logs.ZapGormLog{}, // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond * 20, // Slow SQL threshold
			LogLevel:                  logger.Info,           // Log level
			IgnoreRecordNotFoundError: false,                 // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,                 // Don't include params in the SQL log
			Colorful:                  true,                  // Disable color
		})

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger:                                   logger,
		DryRun:                                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 不转复数
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func InitDB(conn string) {
	DB = NewGorm(conn)
}

// func UseOpentracingPlugin(db *gorm.DB) {
// 	db.Use(&OpentracingPlugin{})
// }
