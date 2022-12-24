package database

import (
	"github.com/dev-hyunsang/clone-twitter-backend/config"
	"github.com/dev-hyunsang/clone-twitter-backend/ent"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL() (*ent.Client, error) {
	client, err := ent.Open("mysql", config.GetEnv("MYSQL"))
	if err != nil {
		return nil, err
	}

	return client, nil
}
