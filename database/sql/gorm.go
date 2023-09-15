package sql

import (
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

type zapLog struct {
}

func (l zapLog) Printf(s string, params ...interface{}) {
	zap.S().Infof(s, params...)
}

func NewGorm(conn string) *gorm.DB {
	var _log logger.Writer
	if zap.S() != nil {
		_log = &zapLog{}
	} else {
		_log = log.New(os.Stdout, "\r\n", log.LstdFlags)
	}

	logger := logger.New(_log, // io writer
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
