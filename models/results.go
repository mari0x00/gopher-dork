package models

import (
	"database/sql"
	"fmt"
)

type Result struct {
	Id          int
	ConfigName  string
	Url         string
	Description string
}

type ResultsService struct {
	DB *sql.DB
}

func (rs *ResultsService) GetAll() ([]Result, error) {
	rows, err := rs.DB.Query(`
		SELECT results.id, configs.name, results.url, results.description
		FROM results 
		INNER JOIN configs 
		ON results.config_id = configs.id;
	`)
	if err != nil {
		return nil, fmt.Errorf("getAll: %w", err)
	}
	results := []Result{}
	for rows.Next() {
		var result Result
		err = rows.Scan(&result.Id, &result.ConfigName, &result.Url, &result.Description)
		if err != nil {
			return nil, fmt.Errorf("getAll: %w", err)
		}
		results = append(results, result)
	}
	return results, nil
}

func (rs *ResultsService) Add(configId int, result Result) error {
	sqlStatement := `
		INSERT INTO results (config_id, url, description)
		VALUES ($1, $2, $3)`
	_, err := rs.DB.Exec(sqlStatement, configId, result.Url, result.Description)
	if err != nil {
		return fmt.Errorf("add: %w", err)
	}
	return nil
}

func (rs *ResultsService) GetConfigById(configId int) (*Config, error) {
	row := rs.DB.QueryRow(`SELECT name, query, results_limit, proxy FROM configs WHERE id = $1`, configId)
	var config Config
	err := row.Scan(&config.Name, &config.Query, &config.Limit, &config.Proxy)
	if err != nil {
		return nil, fmt.Errorf("getConfigById: %w", err)
	}
	return &config, nil
}
