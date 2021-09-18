package test

import (
	"errors"
	"gfeder"
	"gfeder/session"
	"testing"
)

func TestTransaction(t *testing.T) {
	t.Run("rollback", func(t *testing.T) {
		transactionRollback(t)
	})

	t.Run("commit", func(t *testing.T) {
		transactionCommit(t)
	})
}

func transactionRollback(t *testing.T) {
	e, _ := gfeder.NewEngine("sqlite3", "feder.db")
	defer e.Close()
	s := e.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := e.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return nil, errors.New("Error")
	})

	if nil == err || s.HasTable() {
		t.Fatal("failed to rollback")
	}
}

func transactionCommit(t *testing.T) {
	e, _ := gfeder.NewEngine("sqlite3", "feder.db")
	defer e.Close()
	s := e.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := e.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return
	})

	u := &User{}
	s.First(u)

	if err != nil || u.Name != "Tom" {
		t.Fatal("failed to commit")
	}
}
