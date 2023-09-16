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

type Message struct {
	Msg string
	Err string
}

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
	rs := controllers.Results{
		ResultsService: resultsService,
	}
	cs := controllers.Configs{
		ConfigsService: configsService,
	}

	rs.Templates.GetAll = views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "home.gohtml"))
	cs.Templates.GetAll = views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "config.gohtml"))

	// Setup routes
	r.Get("/", rs.GetAll)
	r.Get("/config", cs.GetAll)
	r.Post("/config/add", cs.Add)
	r.Get("/config/delete/{id}", cs.Delete)
	r.Get("/run/{id}", rs.GetDorks)
	r.Get("/run", rs.RunAll)

	fmt.Println("Server started!")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
