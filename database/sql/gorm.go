package sql

import (
	"fmt"
	"time"

	"github.com/Ho-J/base/config"
	"github.com/Ho-J/base/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DBM map[string]*gorm.DB

func NewGorm(mysqls config.Mysqls) map[string]*gorm.DB {
	var dbM = map[string]*gorm.DB{}
	for i, conf := range mysqls.Mysqls {
		if conf.Dbname == "" {
			panic("Dbname 不能为空")
		}
		logger := logger.New(logs.ZapGormLog{}, // io writer
			logger.Config{
				SlowThreshold:             time.Millisecond * 20, // Slow SQL threshold
				LogLevel:                  logger.Info,           // Log level
				IgnoreRecordNotFoundError: false,                 // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,                 // Don't include params in the SQL log
				Colorful:                  true,                  // Disable color
			})

		db, err := gorm.Open(mysql.Open(conf.Conn), &gorm.Config{
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

		dbM[conf.Dbname] = db

		if i == 0 {
			dbM["default"] = db
		}
	}

	return dbM
}

func InitDB(mysqlC config.Mysqls) {
	DBM = NewGorm(mysqlC)
	fmt.Printf("%v", DBM)
}

func GetDb(dbName ...string) *gorm.DB {
	if len(dbName) == 0 {
		return DBM["default"]
	}

	return DBM[dbName[0]]
}
