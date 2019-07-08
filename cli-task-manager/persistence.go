package main

import (
	"github.com/boltdb/bolt"
	"log"
)

var dbService *bolt.DB

func init() {
	db, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		log.Fatal("error during opening db connection:", err)
	}
	dbService = db
}

func save(task *task) error {
	if err := dbService.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("tasks2"))
		if err != nil {
			log.Println("error during creating bucket")
		}
		err = b.Put([]byte(task.description), []byte(task.status))
		return err
	}); err != nil {
		log.Println("error during saving task with description", task.description)
		return err
	}

	log.Println("task saved in db")
	return nil
}

func getAll() []task {
	var out []task
	var task task

	if err := dbService.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks2"))
		log.Println("0")
		if err := b.ForEach(func(k, v []byte) error {
			log.Println("1")
			task.description = string(k)
			task.status = string(v)
			log.Println("2")
			out = append(out, task)
			log.Println("3")
			return nil
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Println("error during creating bolt view", err)
	}

	return out
}
