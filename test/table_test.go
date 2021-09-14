package test

import (
	"gfeder"
	"gfeder/dialect"
	"gfeder/schema"
	"testing"
)

var TestDial, _ = dialect.GetDialect("sqlite3")

type User struct {
	Name string `feder:"PRIMARY KEY"`
	Age  int
}

func TestParse(t *testing.T) {
	schema := schema.Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}

func TestSession_CreateTable(t *testing.T) {
	e, _ := gfeder.NewEngine("sqlite3", "feder.db")
	s := e.NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
