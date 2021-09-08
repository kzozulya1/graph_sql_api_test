package storage

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/go-pg/pg/v10"
)

// dbLogger  struct for DB conn query logging
type dbLogger struct {
	Database string
}

//BeforeQuery logs SQL queries
func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	var fmtQuery, err = q.FormattedQuery()
	if err != nil {
		return c, fmt.Errorf("error from q.FormattedQuery: %s", err.Error())
	}

	logrus.Infof("SQL debug (db %s): %s\n", d.Database, string(fmtQuery))
	return c, nil
}

//AfterQuery for after SQL features
func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	return nil
}

// initDB initialized SB
func InitDB(dbConn string, logSQL bool) (*pg.DB, error) {
	opts, err := pg.ParseURL(dbConn)
	if err != nil {
		return nil, fmt.Errorf("pg.ParseURL %s err: %s", dbConn, err.Error())
	}

	var db = pg.Connect(opts)
	//Add log SQL hook
	if logSQL {
		db.AddQueryHook(dbLogger{
			Database: opts.Database,
		})
	}

	if err := db.Ping(context.Background()); err != nil {
		return nil, err
	}
	return db, err
}
