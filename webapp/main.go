package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type Todo struct {
	bun.BaseModel `bun:"table:todos,alias:t"`

	ID        int64     `bun:"id,pk,autoincrement"`
	Content   string    `bun:"content,notnull"`
	Done      bool      `bun:"done"`
	Until     time.Time `bun:"until,nullzero"`
	CreatedAt time.Time
	UpdatedAt time.Time `bun:",nullzero"`
	DeletedAt time.Time `bun:",soft_deleted,nullzero"`
}

func main() {

	var (
		host     = os.Getenv("POSTGRES_HOSTNAME")
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	databaseUrl := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", user, password, host, dbname)

	sqldb, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer sqldb.Close()

	db := bun.NewDB(sqldb, pgdialect.New())

	defer db.Close()

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.FromEnv("BUNDEBUG"),
	))

	ctx := context.Background()

	_, err = db.NewCreateTable().Model((*Todo)(nil)).IfNotExists().Exec(ctx)

	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})

	e.Logger.Fatal(e.Start(":8989"))
}
