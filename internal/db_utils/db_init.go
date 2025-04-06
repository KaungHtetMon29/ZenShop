package db_utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  string
}

func NewDBConfig(config DBConfig) *DBConfig {
	return &DBConfig{
		Host:     config.Host,
		Port:     config.Port,
		User:     config.User,
		Password: config.Password,
		DbName:   config.DbName,
		SSLMode:  config.SSLMode,
	}
}	

func (dbConfig *DBConfig) GetDSN() string {
	return "host=" + dbConfig.Host +
		" port=" + dbConfig.Port +
		" user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.DbName +
		" sslmode=" + dbConfig.SSLMode
}
func (dbConfig *DBConfig) GetDSNWithTimeZone(timeZone string) string {
	return "host=" + dbConfig.Host +
		" port=" + dbConfig.Port +
		" user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.DbName +
		" sslmode=" + dbConfig.SSLMode +
		" TimeZone=" + timeZone
}
func (dbConfig *DBConfig) ConnectDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}