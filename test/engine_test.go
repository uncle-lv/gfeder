package test

import (
	"fmt"
	"gfeder"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestEngine(t *testing.T) {
	engine, _ := gfeder.NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Jerry").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Test success, %d Row(s) has/have affected\n", count)
}
