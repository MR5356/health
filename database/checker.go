package database

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/MR5356/health"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

const (
	DBDriverMySQL      = "mysql"
	DBDriverSQLite     = "sqlite3"
	DBDriverPostgreSQL = "postgres"

	VersionUnknown = "unknown"
)

var dbPool = map[string]*sql.DB{}

type Checker struct {
	dbDsn        string
	dbDriverType string
}

type Result struct {
	Version string `json:"version"`
	Error   error  `json:"error"`
}

func NewChecker(dbDriverType, dsn string) *Checker {
	return &Checker{
		dbDsn:        dsn,
		dbDriverType: dbDriverType,
	}
}

func (c *Checker) getSqlVersionSQL() string {
	switch c.dbDriverType {
	case DBDriverSQLite:
		return "SELECT sqlite_version()"
	default:
		return "SELECT VERSION()"
	}
}

func (c *Checker) Check() (result *health.Health) {
	result = health.NewHealth()

	id := fmt.Sprintf("%x", md5.Sum([]byte(c.dbDriverType+c.dbDsn)))

	var db *sql.DB
	var ok bool
	var err error
	// check if db is already open
	if db, ok = dbPool[id]; !ok {
		db, err = sql.Open(c.dbDriverType, c.dbDsn)
		if err != nil {
			result.Down()
			result.SetResult(&Result{
				Error: err,
			})
			return
		}
		dbPool[id] = db
	}

	start := time.Now()
	err = db.Ping()
	rtt := time.Since(start).Milliseconds()
	result.SetRTT(rtt)
	if err != nil {
		result.Down()
		result.SetResult(&Result{
			Error: err,
		})
		return
	}

	result.Up()

	// check version
	var version string
	err = db.QueryRow(c.getSqlVersionSQL()).Scan(&version)
	if err != nil {
		result.SetResult(&Result{
			Error:   err,
			Version: VersionUnknown,
		})
		return
	}
	result.SetResult(&Result{
		Version: version,
	})
	return
}
