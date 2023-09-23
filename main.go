package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/mari0x00/google-dork/controllers"
	"github.com/mari0x00/google-dork/migrations"
	"github.com/mari0x00/google-dork/models"
	"github.com/mari0x00/google-dork/templates"
	"github.com/mari0x00/google-dork/views"
)

type config struct {
	PSQL   models.PostgresConfig
	Server struct {
		Address string
	}
}

func loadEnvConfig() (config, error) {
	var cfg config
	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}

	cfg.PSQL = models.PostgresConfig{
		Host:     os.Getenv("PSQL_HOST"),
		Port:     os.Getenv("PSQL_PORT"),
		User:     os.Getenv("PSQL_USER"),
		Password: os.Getenv("PSQL_PASSWORD"),
		Database: os.Getenv("PSQL_DATABASE"),
		SSLMode:  os.Getenv("PSQL_SSLMODE"),
	}
	if cfg.PSQL.Host == "" && cfg.PSQL.Port == "" {
		return cfg, fmt.Errorf("No PSQL config provided.")
	}

	cfg.Server.Address = os.Getenv("SERVER_ADDRESS")
	return cfg, nil
}

func main() {
	cfg, err := loadEnvConfig()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.StripSlashes)

	db, err := models.Open(cfg.PSQL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// Setup services
	resultsService := &models.ResultsService{
		DB: db,
	}
	configsService := &models.ConfigsService{
		DB: db,
	}

	// Setup controllers
	re := controllers.Results{
		ResultsService: resultsService,
	}
	co := controllers.Configs{
		ConfigsService: configsService,
	}

	re.Templates.GetAll = views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "home.gohtml"))
	co.Templates.GetAll = views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "config.gohtml"))

	// Setup routes
	r.Get("/", re.GetAll)
	r.Get("/config", co.GetAll)
	r.Post("/config/add", co.Add)
	r.Get("/config/delete/{id}", co.Delete)
	r.Get("/config/run/{id}", re.GetDorks)
	r.Get("/run", re.RunAll)
	r.Get("/edit/{id}/{status}", re.ChangeStatus)

	fmt.Printf("Starting the webserver on port %s...\n", cfg.Server.Address)
	err = http.ListenAndServe(cfg.Server.Address, r)
	if err != nil {
		panic(err)
	}
}
