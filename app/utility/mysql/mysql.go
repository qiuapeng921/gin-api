package mysql

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

type Manager struct {
	*xorm.Engine

	conf *Config
}

func NewManager(conf *Config) *Manager {
	return &Manager{conf: conf}
}

func (db *Manager) Connect() error {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		db.conf.Username,
		db.conf.Password,
		db.conf.Host,
		db.conf.Database,
		db.conf.Charset)
	connection, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		return err
	}
	if err := connection.Ping(); err != nil {
		return err
	}
	connection.SetMaxIdleConns(db.conf.MaxIdle)
	connection.SetMaxOpenConns(db.conf.MaxOpen)
	show := os.Getenv("APP_ENV")
	if show == gin.DebugMode {
		connection.ShowSQL(false)
		connection.Logger().SetLevel(log.LOG_INFO)
	}
	fmt.Println("Connect mysql success:", db.conf.Host)
	db.Engine = connection
	return nil
}
