// webserver
package main

import (
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data []Post, setup []Settings) {

	t, err := template.ParseGlob("templates/*.html")
	//	if err != nil {
	//		log.Fatal("Error loading templates:" + err.Error())
	//		http.Error(w, "error 500:"+" "+err.Error(), http.StatusInternalServerError)
	//	}
	checkErr(err, "Error loading templates")

	if data != nil && setup == nil {
		t.ExecuteTemplate(w, tmpl, data)
	} else if data == nil && setup != nil {
		t.ExecuteTemplate(w, tmpl, setup)
	} else if data == nil && setup == nil {
		t.ExecuteTemplate(w, tmpl, nil)
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "main", nil, nil)

}

func newsHandler(w http.ResponseWriter, r *http.Request) {

	news := getPosts()
	renderTemplate(w, "news", news, nil)

}

func loadNewsHandler(w http.ResponseWriter, r *http.Request) {

	db := initializationDatabase()
	defer db.Close()
	news := searchNews(&db, r.FormValue("find"))
	renderTemplate(w, "news", news, nil)

}
func settingsHandler(w http.ResponseWriter, r *http.Request) {

	db := initializationDatabase()
	defer db.Close()
	setup := getSettings(&db)
	renderTemplate(w, "settings", nil, setup)

}

func saveSettingsHandler(w http.ResponseWriter, r *http.Request) {
	var setup Settings
	setup.Host = r.FormValue("host")
	setup.Port = r.FormValue("port")
	setup.Url = r.FormValue("url")
	setup.Rule.Item = r.FormValue("item")
	setup.Rule.Title = r.FormValue("title")
	setup.Rule.Text = r.FormValue("text")

	db := initializationDatabase()
	defer db.Close()
	updateSettings(&db, setup)
	http.Redirect(w, r, "/", 302)
}

func startServer(host, port string) {
	log.Println("Listening on port :" + PORT)
	log.Println("Goto http://" + HOST + ":" + PORT)
	log.Println("Parsing site: " + URL)
	log.Printf("Parsing settings: \n %s \n %s \n %s  ", RULES.Item, RULES.Title, RULES.Text)

	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/news", newsHandler)
	http.HandleFunc("/settings", settingsHandler)
	http.HandleFunc("/save", saveSettingsHandler)
	http.HandleFunc("/load", loadNewsHandler)

	log.Println(http.ListenAndServe(host+":"+port, nil))
}
