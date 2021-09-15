package test

import (
	"gfeder/clause"
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {
	var cla clause.Clause
	cla.Set(clause.LIMIT, 3)
	cla.Set(clause.SELECT, "User", []string{"*"})
	cla.Set(clause.WHERE, "Name = ?", "Tom")
	cla.Set(clause.ORDERBY, "Age ASC")
	sql, vars := cla.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	t.Log(sql, vars)
	if sql != "SELECT * FROM User WHERE Name = ? ORDER BY Age ASC LIMIT ?" {
		t.Fatal("failed to build SQL")
	}
	if !reflect.DeepEqual(vars, []interface{}{"Tom", 3}) {
		t.Fatal("failed to build SQLVars")
	}
}

func TestClause_Build(t *testing.T) {
	t.Run("select", func(t *testing.T) {
		TestSelect(t)
	})
}
