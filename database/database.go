package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	// lets sql package be aware of this driver
	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() (*sqlx.DB, error) {
    // exactly the same as the built-in
    db, err := sqlx.Open("sqlite3", "sqlite.db")
    // force a connection and test that it worked
    if err != nil {
        return nil, err
    }
    // check that the database if reachable.
    if err = db.Ping(); err != nil {
        return nil, err 
    }
    return db, nil
}

func Migrate() error {
    db, err := OpenDatabase()

    if err != nil {
       return fmt.Errorf("cannot open the database: %w", err)
    }

    if err = createTodoTable(db); err != nil {
        return fmt.Errorf("cannot create user table: %w", err)
    }

    return nil
}

func createTodoTable(db *sqlx.DB)  error {
    var query string = `CREATE TABLE IF NOT EXISTS user (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username VARCHAR(255),
        pronouns VARCHAR(255),
        bio VARCHAR(255),
        avatar VARCHAR(255),
        theme VARCHAR(255),
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        last_login DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    `
    if _, err := db.Exec(query); err != nil {
        return err
    }
    return nil
}

