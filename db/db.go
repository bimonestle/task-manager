package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var completeBucket = []byte("completedTasks")
var db *bolt.DB

// Data structurw of Task{Key int, Value string}
type Task struct {
	Key   int
	Value string
}

// Initialize or Connect the database
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	fn := func(tx *bolt.Tx) error {
		// _, err := tx.CreateBucketIfNotExists(taskBucket)
		tx.CreateBucketIfNotExists(taskBucket)
		tx.CreateBucketIfNotExists(completeBucket)
		return err
	}
	return db.Update(fn)
}

// Create task and update it to db
func CreateTask(task string) (int, error) {
	var id int

	// Write / Update db to create task
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket) // Get the Bucket
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

// Read all tasks created in db
func AllTasks() ([]Task, error) {
	var tasks []Task

	// View the datas inside db
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket) // Get the Bucket

		// Iterating on the Task keys inside db
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// Read all completed tasks created in db
func AllCompleted() ([]Task, error) {
	var compTasks []Task

	// View the completed tasks
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(completeBucket) // Get the Complete Bucket

		// Iterating the Task keys
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			compTasks = append(compTasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return compTasks, nil
}

// Delete task in db
func DeleteTask(key int) error {
	// Write / Update the db to delete task
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket) // Get the Bucket
		return b.Delete(itob(key))
	})
	return err
}

// Convert Integer to Byte slice
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// Convert Byte slice to Integer
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
