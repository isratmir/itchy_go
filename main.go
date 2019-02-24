package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type api struct {
	DB *sqlx.DB
}

type question struct {
	ID      int    `db:"id" json:"id"`
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
}

func newAPI() (*api, error) {
	api := &api{}
	conn, err := sqlx.Connect("mysql", "root:root@tcp(db:3306)/itchygo")
	if err != nil {
		panic(err)
	}
	api.DB = conn
	return api, nil
}

func (api *api) listQuestions(c echo.Context) error {
	qstns := []question{}
	err := api.DB.Select(&qstns, "SELECT * FROM questions")
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, qstns)
}

func (api *api) insertQuestion(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")
	res, err := api.DB.Exec(
		`INSERT INTO questions (title, content) VALUES (?, ?)`,
		title,
		content,
	)

	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	// return res.LastInsertId(), nil
	return c.JSON(http.StatusOK, id)
}

func main() {
	api, _ := newAPI()
	e := echo.New()
	e.GET("/questions", api.listQuestions)
	e.POST("/questions", api.insertQuestion)
	e.Logger.Fatal(e.Start(":8080"))
}
