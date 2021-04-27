package drivers

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB to holds and store SLQ driver after we make database connection success
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxDBConn = 10
const maxDBIdleConn = 5
const maxDBLifeTime = 5 * time.Minute

// Connect database by specify driver this will pass to Newdatabase function
// driver is depend on what is your database like postgres driver = "pgx"
func ConnectSQL(driver, dsn string) (*DB, error) {
	db, err := NewDatabase(driver, dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(maxDBConn)
	db.SetMaxIdleConns(maxDBIdleConn)
	db.SetConnMaxLifetime(maxDBLifeTime)

	dbConn.SQL = db

	err = testDB(db)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

// Newdatabase to create new database connecttion by specify database driver
// driver is depend on what is your database like postgres driver = "pgx"
func NewDatabase(driver, dsn string) (*sql.DB, error) {
	// create connection to database by using
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// testDB is tries to ping database for connection testing
func testDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	return nil
}
