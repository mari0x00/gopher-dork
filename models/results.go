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
	Status      int
}

type ResultsService struct {
	DB *sql.DB
}

func (rs *ResultsService) GetAll() ([]Result, error) {
	rows, err := rs.DB.Query(`
		SELECT results.id, configs.name, results.url, results.description, results.status
		FROM results 
		INNER JOIN configs 
		ON results.config_id = configs.id
		ORDER BY results.id;
	`)
	if err != nil {
		return nil, fmt.Errorf("getAll: %w", err)
	}
	results := []Result{}
	for rows.Next() {
		var result Result
		err = rows.Scan(&result.Id, &result.ConfigName, &result.Url, &result.Description, &result.Status)
		if err != nil {
			return nil, fmt.Errorf("getAll: %w", err)
		}
		results = append(results, result)
	}
	return results, nil
}

func (rs *ResultsService) Add(configId int, result Result) error {
	sqlStatement := `
		INSERT INTO results (config_id, url, description, status)
		VALUES ($1, $2, $3, 0)`
	_, err := rs.DB.Exec(sqlStatement, configId, result.Url, result.Description)
	if err != nil {
		return fmt.Errorf("add: %w", err)
	}
	return nil
}

func (rs *ResultsService) GetConfigById(configId int) (*Config, error) {
	row := rs.DB.QueryRow(`SELECT name, query, results_limit FROM configs WHERE id = $1`, configId)
	var config Config
	err := row.Scan(&config.Name, &config.Query, &config.Limit)
	if err != nil {
		return nil, fmt.Errorf("getConfigById: %w", err)
	}
	return &config, nil
}

func (rs *ResultsService) GetAllConfigIds() ([]Config, error) {
	rows, err := rs.DB.Query(`SELECT id, query, results_limit FROM configs;`)
	if err != nil {
		return nil, fmt.Errorf("getAllConfigIds: %w", err)
	}
	var configs []Config
	for rows.Next() {
		var entry Config
		err := rows.Scan(&entry.Id, &entry.Query, &entry.Limit)
		if err != nil {
			return nil, fmt.Errorf("getAllConfigIds: %w", err)
		}
		configs = append(configs, entry)
	}
	return configs, nil
}

func (rs *ResultsService) ChangeStatus(recordId int, newStatus int) (int64, error) {
	sqlStatement := "UPDATE results SET status = $1 WHERE id = $2"
	res, err := rs.DB.Exec(sqlStatement, newStatus, recordId)
	if err != nil {
		return 0, fmt.Errorf("changeStatus: %v", err)
	}
	i, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("changeStatus: %v", err)
		return 0, fmt.Errorf("changeStatus: %v", err)
	}
	return i, nil
}

func (rs *ResultsService) DeleteRecord(recordId int) error {
	sqlStatement := "DELETE FROM results WHERE id = $1"
	_, err := rs.DB.Exec(sqlStatement, recordId)
	if err != nil {
		return fmt.Errorf("deleteRecord: %v", err)
	}
	return nil
}
