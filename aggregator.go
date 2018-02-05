package main

import (
	"log"
	"os"
)

type Rules struct {
	Item  string
	Title string
	Text  string
}
type Settings struct {
	Rule Rules
	Url  string
	Host string
	Port string
}
type DbParam struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SslMode  string
	Rebuild  bool
	Setting  bool
}

var (
	URL string = "http://www.e1.ru/news/spool/section_id-105.html " //Url string

	RULES Rules   = Rules{".e1-article-section__content .e1-article", ".e1-article__tit", ".e1-article__text"} // Parse rules
	DB    DbParam = DbParam{"localhost", "5432", "postgres", "test", "aggregator", "disable", false, false}    // Database parametrs
	HOST  string  = "localhost"                                                                                // Server host
	PORT  string  = "8080"                                                                                     // Server port
)

// Checking dfor errors
func checkErr(err error, str string) {
	if err != nil {
		log.Printf("%s", str)
		panic(err)
	}
}

func main() {
	argParser()
	log.SetOutput(os.Stdout)
	//	db, err := sql.Open("postgres", "host="+DB.Host+" port="+DB.Port+" user="+DB.User+" password="+DB.Password+" dbname="+DB.Name+" sslmode="+DB.SslMode)
	//	defer db.Close()

	//	if err != nil {
	//		fmt.Printf("Database opening error -->%v\n", err)
	//		panic("Database error")
	//	}
	db := initializationDatabase()
	defer db.Close()

	if DB.Rebuild {
		rebuildDatabase(&db)
	}

	if DB.Setting {
		setup := getSettings(&db)
		HOST = setup[0].Host
		PORT = setup[0].Port
		URL = setup[0].Url
		RULES = setup[0].Rule
	}

	//	setup := getSettings(&db)
	//	fmt.Printf("Host %s", setup.Host)
	log.Printf("Start server \n")

	startServer(HOST, PORT)

}
