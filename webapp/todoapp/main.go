package main

import (
	"context"
	"database/sql"
	"errors"
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
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
}

type Data struct {
	Todos  []Todo
	Errors []error
}

var (
	host     = os.Getenv("POSTGRES_HOSTNAME")
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_DB")
)

func customFunc(todo *Todo) func([]string) []error {
	return func(values []string) []error {
		if len(values) == 0 || values[0] == "" {
			return nil
		}

		dt, err := time.Parse("2006-01-02T15:04 MST", values[0]+" JST")

		if err != nil {
			return []error{echo.NewBindingError("until", values[0:1], "failed to decode time ", err)}
		}

		todo.Until = dt

		return nil
	}
}

func main() {

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

		var todos []Todo

		ctx := context.Background()

		err := db.NewSelect().Model(&todos).Order("created_at").Scan(ctx)

		fmt.Println(todos)

		if err != nil {
			e.Logger.Error(err)
			return c.Render(http.StatusBadRequest, "index", Data{
				Errors: []error{errors.New("Cannot get todos")},
			})
		}

		return c.String(http.StatusOK, "Hello woraald")
	})

	e.POST("/", func(c echo.Context) error {
		var todo Todo

		errs := echo.FormFieldBinder(c).Int64("id", &todo.ID).String("content", &todo.Content).Bool("done", &todo.Done).CustomFunc("until", customFunc(&todo)).BindErrors()

		if err != nil {
			e.Logger.Error(err)
			return c.Render(http.StatusBadRequest, "index", Data{Errors: errs})
		} else if todo.ID == 0 {
			// idが0の時は登録
			ctx := context.Background()

			if todo.Content == "" {
				err = errors.New("TODO no found")
			} else {
				_, err := db.NewInsert().Model(&todo).Exec(ctx)

				if err != nil {
					e.Logger.Error(err)
					err = errors.New("Canot update")
				}
			}
		} else {
			ctx := context.Background()

			if c.FormValue("delete") != "" {
				// 削除
				_, err = db.NewDelete().Model(&todo).Where("id = ?", todo.ID).Exec(ctx)
			} else {
				// 更新

				var orig Todo

				err = db.NewSelect().Model(&orig).Where("id = ?", todo.ID).Scan(ctx)

				if err != nil {
					orig.Done = todo.Done
					_, err = db.NewUpdate().Model(&orig).Where("id = ?", todo.ID).Exec(ctx)

					if err != nil {
						e.Logger.Error(err)
						err = errors.New("Cannot update")
					}
				}

				if err != nil {
					return c.Render(http.StatusBadRequest, "index", Data{Errors: []error{err}})
				}

			}

		}
		return c.Redirect(http.StatusFound, "/")

	})

	e.Logger.Fatal(e.Start(":8989"))
}
