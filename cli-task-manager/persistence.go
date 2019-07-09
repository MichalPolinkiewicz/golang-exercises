package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	errScanningDbRecord    = "error during scanning db record"
	errQueryExecution      = "error during executing query"
	errTransactionStart    = "error during starting transaction"
	errTransactionCommit   = "error during transaction commit"
	errTransactionRollback = "error during transaction rollback"
)

var dbInstance *sql.DB

func init() {
	dbUri := createDbUri()
	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		log.Println("error during opening db connection:", err)
	}
	dbInstance = db
}

func createDbUri() string {
	e := godotenv.Load()
	if e != nil {
		log.Println("error during loading env file:", e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	charset := os.Getenv("charset")
	pt := os.Getenv("parseTime")
	dbUri := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=%s", username, password, dbName, charset, pt)

	return dbUri
}

func save(t *task) error {
	tx, err := dbInstance.Begin()
	if err != nil {
		log.Println(errTransactionStart, err)
	}

	_, err = tx.Exec("INSERT INTO tasks (name, description, status) VALUES (?, ?, ?)", t.name, t.description, t.status)
	if err != nil {
		log.Println(errQueryExecution, err)
		if tx.Rollback() != nil {
			log.Println(errTransactionRollback, err)
		}
	}

	if tx.Commit() != nil {
		log.Println(errTransactionCommit, err)
	}

	return err
}

func getAll() []task {
	rs, err := dbInstance.Query("SELECT * FROM tasks")
	if err != nil {
		log.Println(errQueryExecution, err)
	}

	var t task
	var tasks []task
	for rs.Next() {
		err := rs.Scan(&t.id, &t.name, &t.description, &t.status)
		if err != nil {
			log.Println(errScanningDbRecord, err)
		} else {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

func getAllByStatus(status *string) []task {
	rs, err := dbInstance.Query("SELECT * FROM tasks WHERE status = ?", status)
	if err != nil {
		log.Println(errQueryExecution, err)
	}

	var t task
	var tasks []task
	for rs.Next() {
		err := rs.Scan(&t.id, &t.name, &t.description, &t.status)
		if err != nil {
			log.Println(errScanningDbRecord, err)
		} else {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

func update(t *task) error {
	tx, err := dbInstance.Begin()
	if err != nil {
		log.Println(errTransactionStart, err)
	}

	rw, err := tx.Exec("UPDATE tasks SET name = ?, description = ? , status = ? WHERE ID = ?", t.name, t.description, t.status, t.id)
	if err != nil {
		log.Println(errQueryExecution, err)
		if tx.Rollback() != nil {
			log.Println(errTransactionRollback, err)
		}
	}

	if tx.Commit() != nil {
		log.Println(errTransactionCommit, err)
	}

	i, err := rw.RowsAffected()
	if int(i) == 0 || err != nil {
		err = errors.New(errQueryExecution)
	}
	return err
}

func getById(id *string) task {
	rs, err := dbInstance.Query("SELECT * FROM tasks WHERE id = ?", id)
	if err != nil {
		log.Println(errQueryExecution, err)
	}

	var t task
	for rs.Next() {
		err := rs.Scan(&t.id, &t.name, &t.description, &t.status)
		if err != nil {
			log.Println(errScanningDbRecord, err)
		}
	}
	return t
}
