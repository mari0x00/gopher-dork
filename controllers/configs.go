package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mari0x00/google-dork/models"
)

type Configs struct {
	Templates struct {
		GetAll Template
	}
	ConfigsService *models.ConfigsService
}

func (c Configs) GetAll(w http.ResponseWriter, r *http.Request) {
	configs, err := c.ConfigsService.GetAll()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	c.Templates.GetAll.Execute(w, configs)
}

func (c Configs) Add(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	query := r.FormValue("query")
	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Limit parameter must be numeric.", http.StatusInternalServerError)
		return
	}
	err = c.ConfigsService.Add(name, query, limit)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/config", http.StatusFound)
	c.GetAll(w, r)
}

func (c Configs) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Limit parameter must be numeric.", http.StatusInternalServerError)
		return
	}
	err = c.ConfigsService.Delete(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	c.GetAll(w, r)
}
