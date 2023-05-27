package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"

	_ "github.com/mattn/go-sqlite3"
)

func showAuthors(db *sql.DB) error {

	rows, err := db.Query(`
	SELECT 
		a.author_id, 
		a.author
	FROM 
	 	autors a
	ORDER BY 
		CAST(a.author_id AS INTEGER)
	`)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var authorID, author string
		err = rows.Scan(&authorID, &author)
		if err != nil {
			return err
		}
		fmt.Printf("%s %s\n", authorID, author)
	}

	return nil
}

func showContent(db *sql.DB, authorId string, titleId string) error {
	var content string

	err := db.QueryRow(`
	SELECT 
	 	c.content
	FROM 
		content c
	WHERE
		c.author_id = ?
	AND c.title_id = ?
	`, authorId, titleId).Scan(&content)

	if err != nil {
		return err
	}

	fmt.Println(content)

	return nil
}

func queryContent(db *sql.DB, query string) error {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())

	if err != nil {
		return err
	}

	seg := t.Wakati(query)

	rows, err := db.Query(`
	SELECT
		a.author_id, 
		a.author, 
		c.title_id, 
		c.title
	FROM 
	    contents c
	INNER JOIN authors a
		ON a.author_id = c.author_id
	INNER JOIN contents_fts f
		ON c.rowid = f.docid
		AND words MATCH ?
	`, strings.Join(seg, " "))

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var authorID, author string
		var titleID, title string

		err = rows.Scan(&authorID, &author, &titleID, &title)

		if err != nil {
			return err
		}

		fmt.Printf("%s % 5s: %s (%s)\n", authorID, titleID, title, author)
	}

	return nil
}

func main() {
	var dsn string

	flag.StringVar(&dsn, "d", "database.sqlite", "database")

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
