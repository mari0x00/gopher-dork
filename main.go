package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"

	"github.com/mari0x00/google-dork/controllers"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fpath := path.Join("templates", "dorks.gohtml")
	tpl, err := template.ParseFiles(path.Join("templates", "tailwind.gohtml"), fpath)
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	fpath := path.Join("templates", "config.gohtml")
	tpl, err := template.ParseFiles(path.Join("templates", "tailwind.gohtml"), fpath)
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

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
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/config", configHandler)
	http.HandleFunc("/run", runHandler)
	fmt.Println("Server started!")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
