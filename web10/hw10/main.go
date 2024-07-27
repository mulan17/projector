//Створити сервер з REST API для перегляду списку cправ.
// Користувач повинен мати можливості: переглядати список завдань, додати нове завдання , 
// змінити існуюче завдання (наприклад, відмітити виконаним), видалити завдання 


package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	s := NewStorage()

	tasks := TaskResource{
		s: s,
	}

	mux.HandleFunc("/tasks", tasks.HandleTasks)
	mux.HandleFunc("/tasks/", tasks.HandleTaskByID)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Failed to listen and serve: %v\n", err)
	}
}

type TaskResource struct {
	s *Storage
}

func (t *TaskResource) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t.GetAll(w, r)
	case http.MethodPost:
		t.CreateOne(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (t *TaskResource) HandleTaskByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		t.UpdateOne(w, r)
	case http.MethodDelete:
		t.DeleteOne(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (t *TaskResource) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks := t.s.GetAllTasks()

	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		fmt.Printf("Failed to encode: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (t *TaskResource) CreateOne(w http.ResponseWriter, r *http.Request) {
	var task Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		fmt.Printf("Failed to decode: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	task.ID = t.s.CreateOneTask(task)

	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		fmt.Printf("Failed to encode: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (t *TaskResource) UpdateOne(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/tasks/"):]
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("Invalid id param: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var task Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		fmt.Printf("Failed to decode: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	task.ID = taskID
	ok := t.s.UpdateTask(task)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func (t *TaskResource) DeleteOne(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/tasks/"):]
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("Invalid id param: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ok := t.s.DeleteTaskByID(taskID)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
