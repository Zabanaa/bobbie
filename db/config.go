package db

import (
    "time"
    "github.com/boltdb/bolt"
)

var tasksBucket = []byte("tasks")

var db *bolt.DB

type Task struct {
    Key int
    Value string
}

func Init(dbPath string) error {

    var err error
    db, err = bolt.Open(dbPath, 0600, &bolt.Options{
        Timeout: 1 * time.Second,
    })

    if err != nil {
        return err
    }

    return db.Update( func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists(tasksBucket)
        return err
    })
}

// Create Task
func CreateTask(task string) (int, error) {

    // 1. takes a string
    // 2. calls db.Update
    // 3. updated the bucket

}


// Delete Task
func DeleteTask(id int) error {

    // 1. takes an int
    // 2. deletes the task at that index
}

// Complete Task
func CompleteTask(id int) error {
    // 1. takes an int (ID)
    // 2. sets the Done attribute to 1
}

// List All Tasks
func ListCurrentTasks() error {
    // 1. returns the entire database where complete is set to 0
}

// List Complete Tasks
func ListCompleteTasks() error {
    // returns the entire database where complete is set to 1
}

