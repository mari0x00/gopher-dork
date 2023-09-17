package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	fmt.Println("RUNNING GET DORKS")
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
	time.Sleep(5 * time.Second)
	http.Redirect(w, r, "/", http.StatusFound)
}
