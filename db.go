// db
package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func initializationDatabase() *sql.DB {
	db, err := sql.Open("postgres", "host="+DB.Host+" port="+DB.Port+" user="+DB.User+" password="+DB.Password+" dbname="+DB.Name+" sslmode="+DB.SslMode)
	checkErr(err, "Database opening error -->%v\n")
	//	defer db.Close()

	return db
}

func rebuildDatabase(pdb **sql.DB) {

	db := *pdb

	init_db_strings := []string{
		"DROP SCHEMA IF EXISTS aggregator CASCADE;",
		"CREATE SCHEMA aggregator;",

		`CREATE TABLE aggregator.settings(
         id serial,
         host varchar(20),
         port varchar(10),
         url varchar(256),
		 item varchar(256),
		 title varchar(256),
		 text varchar(256));`,

		`CREATE TABLE aggregator.data(
         id serial,
         title varchar(256),
         url varchar(256),
         img varchar(256),
		 text text,
		 UNIQUE(url));`,

		"INSERT INTO aggregator.settings (host, port, url, item, title, text)" +
			"VALUES('" + HOST + "','" + PORT + "','" + URL +
			"','" + RULES.Item + "','" + RULES.Title + "','" + RULES.Text + "')"}

	for _, qstr := range init_db_strings {
		_, err := db.Exec(qstr)
		checkErr(err, "Database init error -->%v\n")

	}

	log.Println("Database rebuilded successfully")
}

func getSettings(pdb **sql.DB) []Settings {
	var settings []Settings
	db := *pdb
	log.Println("# Querying")
	rows, err := db.Query("SELECT host, port, url, item, title, text  FROM aggregator.settings LIMIT 1;")
	checkErr(err, "Error DB Query")

	for rows.Next() {
		var setting Settings
		err = rows.Scan(&setting.Host, &setting.Port, &setting.Url, &setting.Rule.Item, &setting.Rule.Title, &setting.Rule.Text)
		checkErr(err, "Error DB rows scan")
		settings = append(settings, setting)
	}
	//	fmt.Println("host | port | url | item | title | text | ")
	//	fmt.Printf("%3v | %8v | %6v | %6v | %6v | %6v\n", setting.Host, setting.Port, setting.Url, setting.Rule.Item, setting.Rule.Title, setting.Rule.Text)

	return settings

}

func updateSettings(pdb **sql.DB, setup Settings) {
	var query string
	db := *pdb

	log.Println("Update database settings")
	query = "UPDATE aggregator.settings SET " +
		"host='" + setup.Host + "', " +
		"port='" + setup.Port + "', " +
		"url='" + setup.Url + "', " +
		"item='" + setup.Rule.Item + "', " +
		"title='" + setup.Rule.Title + "', " +
		"text='" + setup.Rule.Text + "'" +
		" WHERE id = 1;"
		//	log.Fatal(query)
	_, err := db.Exec(query)
	checkErr(err, "Error DB Query")

	HOST = setup.Host
	PORT = setup.Port
	URL = setup.Url
	RULES = setup.Rule
}

func insertNews(pdb **sql.DB, post Post) {
	var query string
	db := *pdb

	log.Println("Insert parse data")
	query = "INSERT INTO aggregator.data(title, url, img, text)" +
		"VALUES ('" + post.Title + "', '" + post.Url + "', '" + post.Img + "', '" + post.Text + "') ON CONFLICT ON CONSTRAINT data_url_key DO NOTHING;"
	_, err := db.Query(query)
	checkErr(err, "Error DB INSERT Query")

}

func searchNews(pdb **sql.DB, search string) []Post {
	var query string
	var posts []Post
	db := *pdb

	query = "SELECT title, url, img, text  FROM aggregator.data " +
		"WHERE   data.title ILIKE '%" + search + "%' OR " +
		"data.text ILIKE '%" + search + "%'" +
		"ORDER BY  data.title ASC;"

	log.Println("Search string " + query)
	rows, err := db.Query(query)
	checkErr(err, "Error DB SELECT Query")

	for rows.Next() {
		var news Post
		err = rows.Scan(&news.Title, &news.Url, &news.Img, &news.Text)
		checkErr(err, "Error DB rows scan")
		posts = append(posts, news)
	}

	return posts

}
