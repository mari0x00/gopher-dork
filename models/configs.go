package models

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Id    int
	Name  string
	Query string
	Limit int
}

type ConfigsService struct {
	DB *sql.DB
}

func (cs *ConfigsService) GetAll() ([]Config, error) {
	rows, err := cs.DB.Query(`SELECT * FROM configs;`)
	if err != nil {
		return nil, fmt.Errorf("getAll: %w", err)
	}
	configs := []Config{}
	for rows.Next() {
		var config Config
		err = rows.Scan(&config.Id, &config.Name, &config.Query, &config.Limit)
		if err != nil {
			return nil, fmt.Errorf("getAll: %w", err)
		}
		configs = append(configs, config)
	}
	return configs, nil
}

func (cs *ConfigsService) Add(name string, query string, limit int) error {
	sql_statement := `
		INSERT INTO configs (name, query, results_limit) 
		VALUES ($1, $2, $3)`
	_, err := cs.DB.Exec(sql_statement, name, query, limit)
	if err != nil {
		return fmt.Errorf("add: %w", err)
	}
	return nil
}

func (cs *ConfigsService) Delete(configId int) error {
	sqlStatement := `
		DELETE FROM configs where id = $1`
	_, err := cs.DB.Exec(sqlStatement, configId)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}
