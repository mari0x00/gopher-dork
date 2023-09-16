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
	}
	c.Templates.GetAll.Execute(w, configs)
}

func (c Configs) Add(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	query := r.FormValue("query")
	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		http.Error(w, "Limit parameter must be numeric", http.StatusNotFound)
	}
	proxy := r.FormValue("proxy")
	err = c.ConfigsService.Add(name, query, limit, proxy)
	if err != nil {
		fmt.Println(err)
	}
	c.GetAll(w, r)
}

func (c Configs) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusNotFound)
	}
	err = c.ConfigsService.Delete(id)
	if err != nil {
		fmt.Println(err)
	}
	c.GetAll(w, r)
}
