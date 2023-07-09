package web

import (
	"fmt"
	"net/http"
)

func Start() {
	fmt.Println("Started server at https://localhost:8080")
	_ = http.ListenAndServe(":8080", nil)
}

type newTask struct {
	Title       string `form:"title" json:"title" xml:"title"`
	Description string `form:"description" json:"description" xml:"description"`
	DueDate     string `form:"dueDate" json:"dueDate" xml:"dueDate"`
}
