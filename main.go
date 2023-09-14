package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/mari0x00/google-dork/controllers"
	"github.com/mari0x00/google-dork/templates"
	"github.com/mari0x00/google-dork/views"
)

func runHandler(w http.ResponseWriter, r *http.Request) {
	fpath := path.Join("templates", "run.gohtml")
	tpl, err := template.ParseFiles(path.Join("templates", "tailwind.gohtml"), fpath)
	if err != nil {
		panic(err)
	}

	config := controllers.CreateConfig("", "", "", "", 5, 0)
	res, err := controllers.QueryDorks(query, config)
	if err != nil {
		fmt.Println(err)
	}

	err = tpl.Execute(w, res)
	if err != nil {
		panic(err)
	}
}

var query = `ext:(doc | docx | pdf | xls | xlsx | txt | ps | rtf | odt | sxw | psw | ppt | pps | xml | ppt | pptx) (intext:"Internal - KMD A/S" | intext:"Confidential - KMD A/S" | intext:"Confidential - KMD employees only")`

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "home.gohtml"))))
	r.Get("/config", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "config.gohtml"))))
	r.Post("/config", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "config.gohtml"))))
	r.Get("/run", runHandler)

	fmt.Println("Server started!")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
