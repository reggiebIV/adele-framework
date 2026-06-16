package database

import (
	"database/sql"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Database struct {
	DataType string
	Pool     *sql.DB
	Mongo    *mongo.Client
}

type DataSourceName struct {
	Host          string
	Port          string
	User          string
	Password      string
	DatabaseName  string
	SslMode       string
	IsAtlasHosted bool
}
