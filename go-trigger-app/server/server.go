package main

import (
	"database/sql"
	"go-trigger-app/db"
	"go-trigger-app/web"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	d, err := sql.Open("postgres", dataSource())
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()
	// CORS is enabled only in prod profile
	cors := os.Getenv("profile") == "prod"
	app := web.NewApp(db.NewDB(d), cors)
	err = app.Serve()
	log.Println("Error", err)
}

func dataSource() string {
	host := "localhost"
	pass := "pass"
	if os.Getenv("profile") == "prod" {
		host = "db"
		pass = os.Getenv("db_pass")
	}
	return "postgresql://" + host + ":5433/goxygen" +
		"?user=goxygen&sslmode=disable&password=" + pass
}
