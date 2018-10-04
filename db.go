package db

import (
	"fmt"
	"sync"

	//Mysql driver for gorm
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var (
	// Username is a username used for connection
	Username string

	// Password is a password used for connection
	Password string

	// Database is a db name to which you want to connect
	Database string

	// Host for connection
	Host = "localhost"

	// Port used for connection
	Port = 3306

	//Charset for the connection.
	Charset = "utf8"

	// ParseTime defines if we need to parse time columns.
	ParseTime = "True"

	//Location used for timestamps related functions.
	Location = "Europe%2FLondon"

	//DebugLog for now unused
	DebugLog bool
)

var instance *gorm.DB
var dbOnce sync.Once

//Returns (or creates) the singleton instance
func Instance() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		instance, err = CreateInstance(Username, Password, Database, Host, Port, DebugLog)
		if err != nil {
			fmt.Sprintf("Can't connect to database [%s]", err)
			os.Exit(1)
		}
	})
	return instance
}

// CreateInstance
func CreateInstance(username, password, dbName, host string, port int, debugLog bool) (*gorm.DB, error) {
	var connectionString = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		username,
		password,
		host,
		port,
		dbName,
		Charset,
		ParseTime,
		Location,
	)

	dbManager, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, errors.New("Cannot connect to database")
	}
	//if debugLog {
	//	dbManager.LogMode(true)
	//
	//	//
	//	gormLogger := &logger2.GormLogger{}
	//	gormLogger.SetLogrus(logger.GetLogger())
	//	dbManager.SetLogger(gormLogger)
	//}

	dbManager.LogMode(debugLog)
	return dbManager, nil
}
