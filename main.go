package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

var conn *sqlx.DB

type Question struct {
	id      int    `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
}

func insertQuestion(c echo.Context) error {
	conn, err := sqlx.Connect("mysql", "root:root@tcp(db:3306)/itchygo")
	if err != nil {
		panic(err)
	}
	title := c.FormValue("title")
	content := c.FormValue("content")
	// insert := `INSERT INTO questions (title, content) VALUES (?, ?)`
	res, err := conn.Exec(
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

	e := echo.New()
	e.POST("/questions", insertQuestion)
	e.Logger.Fatal(e.Start(":8080"))
}
