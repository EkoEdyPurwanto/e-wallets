package manager

import (
	"EEP/e-wallets/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type InfraManager interface {
	Conn() *sql.DB
}

type infraManager struct {
	db  *sql.DB
	cfg *config.Config
}

// Constructor
func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{
		cfg: cfg,
	}
	err := conn.initDb()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (i *infraManager) Conn() *sql.DB {
	return i.db
}

func (i *infraManager) initDb() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("APP_DB_HOST"),
		viper.GetString("APP_DB_PORT"),
		viper.GetString("APP_DB_USER"),
		viper.GetString("APP_DB_PASSWORD"),
		viper.GetString("APP_DB_NAME"),
	)
	fmt.Println(dsn)

	db, err := sql.Open(viper.GetString("APP_DB_DRIVER"), dsn)
	if err != nil {
		return err
	}
	i.db = db
	return nil
}
