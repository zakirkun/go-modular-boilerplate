package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB
	POSGRES_CONFIG = "user=%s password=%s dbname=%s host=%s port=%s sslmode=%s"
	MYSQL_CONFIG   = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

type DBModel struct {
	ServerMode   string `config:"server_mode"`
	Driver       string `config:"db_driver"`
	Host         string `config:"db_host"`
	Port         string `config:"db_port"`
	Name         string `config:"db_name"`
	Username     string `config:"db_username"`
	Password     string `config:"db_password"`
	MaxIdleConn  int    `config:"conn_idle"`
	MaxOpenConn  int    `config:"conn_max"`
	ConnLifeTime int    `config:"conn_lifetime"`
}

func (c *DBModel) OpenDB() (*gorm.DB, *error) {

	var connection gorm.Dialector

	switch c.Driver {
	case "postgres":
		connectionUrl := fmt.Sprintf(POSGRES_CONFIG, c.Username, c.Password, c.Name, c.Host, c.Port, "disable")
		connection = postgres.Open(connectionUrl)
	case "mysql":
		connectionUrl := fmt.Sprintf(MYSQL_CONFIG, c.Username, c.Password, c.Host, c.Port, c.Name)
		connection = mysql.Open(connectionUrl)
	default:
		log.Fatal("No Database Selected!, Please check config.toml")
		os.Exit(1)
	}

	db, err := gorm.Open(connection, &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot Connect to DB With Message %s", err.Error())
		return nil, &err
	}

	conPool, err := db.DB()
	if err != nil {
		log.Fatalf("Cannot Create Connection Pool to DB With Message %s", err.Error())
		return nil, &err
	}

	/** SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	**/
	conPool.SetMaxIdleConns(c.MaxIdleConn)

	/** SetMaxOpenConns sets the maximum number of open connections to the database.
	**/
	conPool.SetMaxOpenConns(c.MaxOpenConn)

	/** SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	**/
	conPool.SetConnMaxLifetime(time.Duration(c.ConnLifeTime) * time.Minute)

	return db, nil
}
