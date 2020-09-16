package models

import (
	"errors"
	"gotask/database"
	"log"
)

type Task struct {
	UID         uint32 `json:"id"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	Closed      bool   `json:"closed"`
}

func CreateTask(task Task) (error, Task) {
	con := database.Connect()
	defer con.Close()
	tx, err := con.Begin()
	if err != nil {
		return err, Task{}
	}
	sql := "insert into tasks (description, user_id) values ($1, $2) returning uid, created_at"
	stmt, err := tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		return err, Task{}
	}
	defer stmt.Close()
	err = stmt.QueryRow(task.Description, task.UserID).Scan(&task.UID, &task.CreatedAt)
	if err != nil {
		tx.Rollback()
		return err, Task{}
	}
	return tx.Commit(), task
}

func GetTasks(user_id string, closed bool) (error, []Task) {
	con := database.Connect()
	defer con.Close()
	sql := `select * from tasks where user_id = $1 and closed= $2 ORDER BY "created_at" desc`
	rs, err := con.Query(sql, user_id, closed)
	if err != nil {
		return err, nil
	}
	defer rs.Close()
	var tasks []Task
	for rs.Next() {
		var task Task
		err := rs.Scan(&task.UID, &task.Description, &task.UserID, &task.Closed, &task.CreatedAt)
		if err != nil {
			log.Fatal(err)
			return err, nil
		}
		tasks = append(tasks, task)
	}
	return nil, tasks
}

func UpdateTask(task Task) error {
	con := database.Connect()
	defer con.Close()
	sql := "update tasks set closed= $1 where uid = $2 and user_id = $3"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Exec(task.Closed, task.UID, task.UserID)
	if err != nil {
		return err
	}
	count, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Task not found")
	}
	return nil
}

func DeleteTask(task Task) error {
	con := database.Connect()
	defer con.Close()
	sql := "delete from tasks where uid = $1 and user_id = $2"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Exec(task.UID, task.UserID)
	if err != nil {
		return err
	}
	count, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Task not found")
	}
	return nil
}
