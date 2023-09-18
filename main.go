package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mari0x00/google-dork/controllers"
	"github.com/mari0x00/google-dork/migrations"
	"github.com/mari0x00/google-dork/models"
	"github.com/mari0x00/google-dork/templates"
	"github.com/mari0x00/google-dork/views"
)

//TODO: DELETE FROM HERE
var query = `ext:(doc | docx | pdf | xls | xlsx | txt | ps | rtf | odt | sxw | psw | ppt | pps | xml | ppt | pptx) (intext:"Internal - KMD A/S" | intext:"Confidential - KMD A/S" | intext:"Confidential - KMD employees only")`

func main() {
	r := chi.NewRouter()
	cfg := models.DefaultPostgressConfig()
	db, err := models.Open(cfg)
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
	r.Get("/run/{id}", re.GetDorks)
	r.Get("/run", re.RunAll)
	r.Get("/edit/{id}/{status}", re.ChangeStatus)
	r.Get("/edit/{id}/delete", re.RunAll)

	fmt.Println("Server started!")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
