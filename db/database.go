
package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"backendapp/config"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	return sql.Open("postgres", dsn)
}
