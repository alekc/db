package db

import (
	"fmt"
	"sync"
	"time"

	//Mysql driver for gorm
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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

	//Maximum open connections which app can have
	MaxOpenConnections = 10

	//
	MaxIdleConnections = 5

	//
	ConnMaxLifeTime = 10
)

var instance *gorm.DB
var dbOnce sync.Once

//Returns (or creates) the singleton instance
func Instance() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		instance, err = CreateInstance(Username, Password, Database, Host, Port, DebugLog)
		if err != nil {
			fmt.Printf("Error on database instance creation. [%s]\n", err)
			if DebugLog {
				fmt.Printf("Username [%s], Password [%s], Db [%s], Host [%s:%d]\n", Username, Password, Database, Host, Port)
			}
			os.Exit(1)
		}
		instance.DB().SetMaxOpenConns(MaxOpenConnections)
		instance.DB().SetMaxIdleConns(MaxIdleConnections)
		instance.DB().SetConnMaxLifetime(time.Duration(ConnMaxLifeTime) * time.Second)
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
		return nil, err
	}

	dbManager.LogMode(debugLog)
	return dbManager, nil
}
