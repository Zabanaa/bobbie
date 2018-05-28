package db

import (
    "time"
    "encoding/binary"
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
func CreateTask(newTask string) (int, error) {

    var id int

    err := db.Update( func(tx *bolt.TX) error {

        bucket := tx.Bucket(tasksBucket) // retrieve the bucket
        id64, _ := bucket.NextSequence() // get the next auto incremented id
        id = int(id64) // cast that id to an int (so that we can access it outside the closure) and return it
        key := itob(id) // set that to the key (after casting it to a byte slice)
        task := []byte(newTask) // cast the newTask as a byte slice
        return bucket.Put(key, task) // insert the record into the database
    })

    if err != nil {
       return -1, err
    }

    return id, nil
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

// Integer to Byte Slice
func itob(value int) []byte {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(value))
    return b
}

func btoi(b []byte) int {
    return int(binary.BigEndian.Uint64(b))
}
