package configs

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

type Config struct {
	DBHost            string
	DBPort            string
	DBUserName        string
	DBPass            string
	DBName            string
	DBMaxIdleConn     int32
	DBMaxOpenConn     int32
	DBMigrationScheme string
	dbMigrationPath   string
}

func NewConfig() (Config, error) {
	var cfg Config
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort = os.Getenv("DB_PORT")
	cfg.DBUserName = os.Getenv("DB_USERNAME")
	cfg.DBPass = os.Getenv("DB_PASS")
	cfg.DBName = os.Getenv("DB_NAME")
	idleConn := os.Getenv("DB_MAX_IDLE_CONNS")
	maxConn := os.Getenv("DB_MAX_OPEN_CONNS")
	idleConnInt, err := strconv.Atoi(idleConn)
	if err != nil {
		return Config{}, err
	}
	maxConnInt, err := strconv.Atoi(maxConn)
	if err != nil {
		return Config{}, err
	}
	cfg.DBMaxIdleConn = int32(idleConnInt)
	cfg.DBMaxOpenConn = int32(maxConnInt)

	cfg.DBMigrationScheme = os.Getenv("DB_MIGRATION_TABLE")
	cfg.dbMigrationPath = os.Getenv("DB_MIGRATION_PATH")
	return cfg, nil
}

func (c *Config) DBUrl() string {
	urlPostgres := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s pool_max_conns=%d",
		c.DBHost, c.DBPort, c.DBUserName, c.DBName, c.DBPass, c.DBMaxOpenConn)
	return urlPostgres
}

func (c *Config) DBMigrationPath() fs.FS {
	return os.DirFS(c.dbMigrationPath)
}
