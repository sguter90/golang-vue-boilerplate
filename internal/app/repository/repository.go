package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

type DbEnvironment struct {
	Dialect    string `yaml:"dialect"`
	DataSource string `yaml:"datasource"`
	Dir        string `yaml:"dir"`
	TableName  string `yaml:"table"`
	SchemaName string `yaml:"schema"`
}

func readDbConfig() (*DbEnvironment, error) {
	dataSource := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") +
		"@(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME")

	if os.Getenv("DB_DIALECT") == "mysql" {
		dataSource += "?parseTime=true"
	}

	env := DbEnvironment{}
	env.Dialect = os.Getenv("DB_DIALECT")
	env.DataSource = dataSource

	return &env, nil
}

func ConnectDb() (*sqlx.DB, string, error) {
	env, err := readDbConfig()
	if err != nil {
		return nil, "", fmt.Errorf("failed to read database config: %s", err)
	}

	db, err := sqlx.Open(env.Dialect, env.DataSource)
	if err != nil {
		return nil, "", fmt.Errorf("cannot connect to database: %s", err)
	}

	return db, env.Dialect, nil

}
