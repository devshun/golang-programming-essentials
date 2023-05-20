package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
)

func showAuthors(db *sql.DB) error {

	return nil
}

func showContent(db *sql.DB, authorId string, titleId string) error {

	return nil
}

func queryContent(db *sql.DB, query string) error {

	return nil
}

func main() {
	var dsn string

	flag.StringVar(&dsn, "d", "database.sqlite", "database")

	flag.Usage = func() {
		fmt.Println("usage")
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	db, err := sql.Open("sqlite3", dsn)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	switch flag.Arg(0) {
	case "authors":
		err = showAuthors(db)
	case "titles":
		if flag.NArg() != 2 {
			flag.Usage()
			os.Exit(2)
		}
		err = showContent(db, flag.Arg(1), flag.Arg(2))

	case "query":
		if flag.NArg() != 2 {
			flag.Usage()
			os.Exit(2)
		}
		err = queryContent(db, flag.Arg(1))

	}

	if err != nil {
		log.Fatal(err)
	}

}
