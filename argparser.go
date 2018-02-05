// argparser
package main

import (
	"flag"
)

// Parse args command line
func argParser() {
	//Rules for parse args

	flag.StringVar(&URL, "u", URL, "URL address news site for aggregation")
	flag.StringVar(&RULES.Item, "ritem", RULES.Item, "Rules item ")
	flag.StringVar(&RULES.Title, "rtitle", RULES.Title, "Rules title ")
	flag.StringVar(&RULES.Text, "rtext", RULES.Text, "Rules text")

	flag.StringVar(&HOST, "h", HOST, "Addres host")
	flag.StringVar(&PORT, "p", PORT, "Listen port")

	flag.StringVar(&DB.Host, "db-host", DB.Host, "Database host ")
	flag.StringVar(&DB.Port, "db-port", DB.Port, "Database port ")
	flag.StringVar(&DB.User, "db-user", DB.User, "Database user ")
	flag.StringVar(&DB.Password, "db-password", DB.Password, "Database password ")
	flag.StringVar(&DB.Name, "db-name", DB.Name, "Database name ")
	flag.StringVar(&DB.SslMode, "db-ssl", DB.SslMode, "Database ssl mode  ")
	flag.BoolVar(&DB.Rebuild, "db-create", DB.Rebuild, "Create new database schema if not exist or rebuild database if flag true, default false ")
	flag.BoolVar(&DB.Setting, "get-settings-on-db", DB.Setting, "Get settings from DB if flag true, default false (settings by default) ")

	//	flag.PrintDefaults()
	//Run args parser
	flag.Parse()

}
