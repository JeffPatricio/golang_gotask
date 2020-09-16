package controllers

import (
	"encoding/json"
	"gotask/models"
	"gotask/utils"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ResponsePostTask struct {
	Task    interface{} `json:"task"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

type ResponseGetTasks struct {
	Tasks   interface{} `json:"tasks"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

type ResponseDeleteTasks struct {
	TaskID  uint32 `json:"task_id"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type ResponseUpdateTask struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	userID := r.Header.Get("UserId")
	if userID == "" {
		utils.ErrorResponse(w, "Invalid User Id Header", http.StatusUnprocessableEntity)
		return
	}
	var task models.Task
	err := json.Unmarshal(body, &task)
	if err != nil {
		utils.ErrorResponse(w, "Invalid input data format", http.StatusUnprocessableEntity)
		return
	}
	task.UserID = userID
	if utils.IsEmpty(task.Description) || utils.IsEmpty(task.UserID) {
		utils.ErrorResponse(w, "You must fill in the description and user id fields", http.StatusUnprocessableEntity)
		return
	}
	err, task = models.CreateTask(task)
	if err != nil {
		utils.ErrorResponse(w, "Internal error when creating task in the database", http.StatusUnprocessableEntity)
		return
	}

	utils.ToJson(w, ResponsePostTask{Message: "Successfully created task", Success: true, Task: task})
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	paramClosed := r.FormValue("closed")
	closed, err := strconv.ParseBool(paramClosed)
	if err != nil {
		utils.ErrorResponse(w, "Error in parameter closed in url", http.StatusBadRequest)
		return
	}
	err, tasks := models.GetTasks(r.Header.Get("UserId"), closed)
	if err != nil {
		utils.ErrorResponse(w, "Internal error when fetching tasks from the database", http.StatusBadRequest)
		return
	}
	if tasks == nil {
		utils.ToJson(w, ResponseGetTasks{Message: "OK", Success: true, Tasks: []models.Task{}})
		return
	}
	utils.ToJson(w, ResponseGetTasks{Message: "OK", Success: true, Tasks: tasks})
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &task)
	if err != nil {
		utils.ErrorResponse(w, "Invalid input data format", http.StatusUnprocessableEntity)
		return
	}
	task.UserID = r.Header.Get("UserId")
	if task.UserID == "" {
		utils.ErrorResponse(w, "Invalid User Id Header", http.StatusUnprocessableEntity)
		return
	}
	params := mux.Vars(r)
	if params["task_id"] == "" {
		utils.ErrorResponse(w, "It is necessary to inform the task id in url parameters", http.StatusUnprocessableEntity)
		return
	}
	taskID64, err := strconv.ParseInt(params["task_id"], 10, 32)
	if err != nil {
		utils.ErrorResponse(w, "Invalid task id param", http.StatusUnprocessableEntity)
		return
	}
	task.UID = uint32(taskID64)

	err = models.UpdateTask(task)

	if err != nil && err.Error() == "Task not found" {
		utils.ErrorResponse(w, "Task not found", http.StatusBadRequest)
		return
	}
	if err != nil {
		utils.ErrorResponse(w, "Internal error when updating a task", http.StatusBadRequest)
		return
	}
	utils.ToJson(w, ResponseUpdateTask{Message: "Task successfully updated", Success: true})
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	task.UserID = r.Header.Get("UserId")
	if task.UserID == "" {
		utils.ErrorResponse(w, "Invalid User Id Header", http.StatusUnprocessableEntity)
		return
	}
	params := mux.Vars(r)
	if params["task_id"] == "" {
		utils.ErrorResponse(w, "It is necessary to inform the task id in url parameters", http.StatusUnprocessableEntity)
		return
	}
	taskID64, err := strconv.ParseInt(params["task_id"], 10, 32)
	if err != nil {
		utils.ErrorResponse(w, "Invalid task id param", http.StatusUnprocessableEntity)
		return
	}
	task.UID = uint32(taskID64)
	err = models.DeleteTask(task)

	if err != nil && err.Error() == "Task not found" {
		utils.ErrorResponse(w, "Task not found", http.StatusNotFound)
		return
	}
	if err != nil {
		utils.ErrorResponse(w, "Internal error when deleting a task", http.StatusBadRequest)
		return
	}
	utils.ToJson(w, ResponseDeleteTasks{Message: "Task successfully deleted", Success: true, TaskID: task.UID})
}
