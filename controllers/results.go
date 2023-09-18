package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgconn"
	dorker "github.com/mari0x00/google-dork/cmd"
	"github.com/mari0x00/google-dork/models"
)

type Results struct {
	Templates struct {
		GetAll Template
		Run    Template
	}
	ResultsService *models.ResultsService
}

func (re Results) GetAll(w http.ResponseWriter, r *http.Request) {
	results, err := re.ResultsService.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	re.Templates.GetAll.Execute(w, results)
}

func (re Results) GetDorks(w http.ResponseWriter, r *http.Request) {
	configId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
	}
	config, err := re.ResultsService.GetConfigById(configId)
	if err != nil {
		fmt.Printf("getDorks: %v\n", err)
	}
	results, err := dorker.Dork(config.Query, config.Limit)
	if err != nil {
		fmt.Printf("getDorks: %v\n", err)
	}
	for _, result := range results {
		entry := models.Result{
			ConfigName:  config.Name,
			Description: result.Name,
			Url:         result.Url,
		}
		err := re.ResultsService.Add(configId, entry)
		if err != nil {
			var duplicateEntryError = &pgconn.PgError{Code: "23505"}
			if errors.As(err, &duplicateEntryError) {
				continue
			}
			fmt.Printf("getDorks: %v\n", err)
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func (re Results) RunAll(w http.ResponseWriter, r *http.Request) {
	configs, err := re.ResultsService.GetAllConfigIds()
	if err != nil {
		fmt.Println(err)
	}
	for _, config := range configs {
		results, err := dorker.Dork(config.Query, config.Limit)
		if err != nil {
			fmt.Printf("getDorks: %v\n", err)
		}
		for _, result := range results {
			entry := models.Result{
				ConfigName:  config.Name,
				Description: result.Name,
				Url:         result.Url,
			}
			err := re.ResultsService.Add(config.Id, entry)
			if err != nil {
				var duplicateEntryError = &pgconn.PgError{Code: "23505"}
				if errors.As(err, &duplicateEntryError) {
					continue
				}
				fmt.Printf("getDorks: %v\n", err)
			}
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func (re Results) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	recordId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	newStatus, err := strconv.Atoi(chi.URLParam(r, "status"))
	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	if newStatus < 0 || newStatus > 2 {
		http.Error(w, "Invalid status code provided", http.StatusNotFound)
		return
	}
	res, err := re.ResultsService.ChangeStatus(recordId, newStatus)
	if err != nil {
		println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
	}
	if res < 1 {
		http.Error(w, "Invalid recordId.", http.StatusNotFound)
		return
	}
	_, err = w.Write([]byte(chi.URLParam(r, "status")))
	if err != nil {
		println(err)
	}
}
