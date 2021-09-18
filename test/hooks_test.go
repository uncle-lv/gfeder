package test

import (
	"gfeder"
	"gfeder/log"
	"gfeder/session"
	"testing"
)

type Account struct {
	ID       int `feder:"PRIMARY KEY"`
	Password string
}

func (account *Account) BeforeInsert(s *session.Session) error {
	log.Info("before inert", account)
	account.ID += 1000
	return nil
}

func (account *Account) AfterQuery(s *session.Session) error {
	log.Info("after query", account)
	account.Password = "******"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	e, _ := gfeder.NewEngine("sqlite3", "feder.db")
	s := e.NewSession().Model(&Account{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, _ = s.Insert(&Account{1, "123456"}, &Account{2, "abcdef"})

	account := &Account{}

	err := s.First(account)
	if err != nil || account.ID != 1001 || account.Password != "******" {
		t.Fatal("Failed to call hooks after query, got", account)
	}
}
