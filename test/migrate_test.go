package test

import (
	"gfeder"
	"reflect"
	"testing"
)

type Product struct {
	Name string `feder:"PRIMARY KEY"`
	Num  int
}

func TestMigrate(t *testing.T) {
	e, _ := gfeder.NewEngine("sqlite3", "feder.db")
	defer e.Close()
	s := e.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE Product(Name text PRIMARY KEY, Count integer);").Exec()
	_, _ = s.Raw("INSERT INTO Product(`Name`) values (?), (?)", "chocolate", 3).Exec()
	e.Migrate(&Product{})
	rows, _ := s.Raw("SELECT * FROM Product").QueryRows()
	columns, _ := rows.Columns()
	if !reflect.DeepEqual(columns, []string{"Name", "Num"}) {
		t.Fatal("Failed to migrate table Product, got columns", columns)
	}
}
