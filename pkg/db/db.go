package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var db *gorm.DB
var once sync.Once

func DB() *gorm.DB {
	return db
}

func Init(cfg Config) {
	once.Do(func() {
		dsn := fmt.Sprintf("%s?charset=%s", cfg.Link, cfg.Charset)
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}

		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(cfg.MaxIdle)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(cfg.MaxOpen)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(cfg.MaxLifetime)

	})

}
