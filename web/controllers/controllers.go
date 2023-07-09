package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-client/service"
)

var client service.TaskAppClient

func StartController(c service.TaskAppClient) {
	client = c
}

var RootController http.Handler = http.FileServer(http.Dir("./web/static"))

func ListTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received list request")
	entries, err := client.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "Failed to get tasks: %v", err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		content, _ := json.Marshal(entries)
		w.Write(content)
	}
}

func CreateTasks(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Form is: %v\n", r)
	}
	title := r.Form.Get("title")
	desc := r.Form.Get("description")
	dueDate := r.Form.Get("dueDate")

	fmt.Printf("Received create request with title: %s, desc: %s, duedate: %s\n", title, desc, dueDate)

	entry, err := client.Create(title, desc, dueDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "Failed to get tasks: %v", err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		content, _ := json.Marshal(entry)
		_, err := w.Write(content)
		if err != nil {
			return
		}
	}

}
