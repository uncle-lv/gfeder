package gfeder

import (
	"database/sql"
	"gfeder/dialect"
	"gfeder/log"
	"gfeder/session"
)

type TxFunc func(*session.Session) (result interface{}, err error)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)

	if err != nil {
		log.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	dialect, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}

	e = &Engine{db: db, dialect: dialect}
	log.Info("Connect database succeed")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database succeed")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}

func (engine *Engine) Transaction(f TxFunc) (result interface{}, err error) {
	s := engine.NewSession()
	if err := s.Begin(); err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = s.Rollback()
			panic(p)
		} else if err != nil {
			_ = s.Rollback()
		} else {
			defer func() {
				if err != nil {
					_ = s.Rollback()
				}
			}()
			err = s.Commit()
		}
	}()

	return f(s)
}
