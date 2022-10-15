package main

type DBConfig struct {
	DbName            string
	DbUser            string
	DbPassword        string
	DbHost            string
	DbPort            int
	Location          string
	DbConnMaxLifetime int
	DbMaxOpenConns    int
	DbMaxIdleConns    int
	DbConnMaxIdleTime int
	DbSuffix          string
}

func NewDBConfig(
	dbName string,
	dbUser string,
	dbPassword string,
	dbHost string,
	dbPort int,
	location string,
	dbConnMaxLifetime int,
	dbMaxOpenConns int,
	dbMaxIdleConns int,
	dbConnMaxIdleTime int,
) *DBConfig {
	return &DBConfig{
		DbName:            dbName,
		DbUser:            dbUser,
		DbPassword:        dbPassword,
		DbHost:            dbHost,
		DbPort:            dbPort,
		Location:          location,
		DbConnMaxLifetime: dbConnMaxLifetime,
		DbConnMaxIdleTime: dbConnMaxIdleTime,
		DbMaxIdleConns:    dbMaxIdleConns,
		DbMaxOpenConns:    dbMaxOpenConns,
	}
}
