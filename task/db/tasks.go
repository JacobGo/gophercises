package db

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
)

var db *bolt.DB
var taskBucket = []byte("tasks")

type Task struct {
	ID int
	Description string
	Done bool
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func Init() error {
	var err error
	db, err = bolt.Open("db/tasks.db", 0600, nil)
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket(taskBucket)
		return err
	})
}

func CreateTask(desc string) (int, error) {
	var id int
	var err error
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		nid, _ := bucket.NextSequence()
		id = int(nid)

		task := Task{
			ID:          id,
			Description: desc,
			Done:        false,
		}

		encoded, err := json.Marshal(task)
		if err != nil {
			return err
		}

		return bucket.Put(itob(id), encoded)
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func ListTasks() ([]Task, error) {
	var tasks []Task
	var err error
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		_ = bucket.ForEach(func(id, encoded []byte) error {
			task := Task{}
			_ = json.Unmarshal(encoded, &task)
			tasks = append(tasks, task)
			return nil
		})
		return nil
	})
	if err != nil {
		return []Task{}, err
	}
	return tasks, nil
}

func MarkTaskDone(taskId int) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		encoded := bucket.Get(itob(taskId))
		if encoded == nil {
			return errors.New("taskId not found")
		}
		task := Task{}
		_ = json.Unmarshal(encoded, &task)
		task.Done = true

		encoded, _ = json.Marshal(task)
		return bucket.Put(itob(taskId), encoded)
	})
}