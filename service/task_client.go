package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

type TaskAppClient struct {
	protoClient TaskProtoClient
	conn        *grpc.ClientConn
}

func StartTaskAppClient() TaskAppClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return TaskAppClient{protoClient: NewTaskProtoClient(conn), conn: conn}
}

func (client TaskAppClient) List() (*[]*Task, error) {
	req := &ListTaskRequest{}
	resp, err := client.protoClient.List(context.Background(), req)
	if err != nil {
		log.Printf("Failed to list tasks: %v", err)
		return nil, err
	}
	log.Printf("There are %d tasks\n", len(resp.Tasks))
	listTasks(resp)
	return &resp.Tasks, nil
}

func (client TaskAppClient) Create(title string, description string, duedate string) (*Task, error) {
	if description == "" {
		description = ""
	}
	if duedate == "" {
		duedate = time.Now().Format("2006-01-02 15:04")
		fmt.Println("No duedate provided, using current time " + duedate)
	}
	req := CreateTaskRequest{Title: title, Description: &description, DueDate: &duedate}
	resp, err := client.protoClient.Create(context.Background(), &req)
	if err != nil {
		log.Printf("Failed to create task: %v with title %s", err, title)
		return nil, err
	} else {
		fmt.Println("Created new task with below details: ")
		fmt.Println(toString(resp.Task))
	}
	return resp.Task, nil
}

func (client TaskAppClient) Details(taskId int32) (*Task, error) {
	req := DetailsTaskRequest{TaskId: taskId}
	resp, err := client.protoClient.Details(context.Background(), &req)
	if err != nil {
		log.Printf("Failed to get details for task id %d with error: %v\n", taskId, err)
		return nil, err
	} else {
		fmt.Println(toString(resp.Task))
	}
	return resp.Task, nil
}

func (client TaskAppClient) UpdateTask(taskId int32,
	newTitle *string,
	newDesc *string,
	newDueDate *string,
	newStatus *TaskStatus) (*Task, error) {
	req := UpdateTaskRequest{TaskId: taskId, NewTitle: newTitle, NewDescription: newDesc, NewDueDate: newDueDate, NewStatus: newStatus}
	resp, err := client.protoClient.Update(context.Background(), &req)
	if err != nil {
		log.Printf("Failed to update task id %d with error: %v\n", taskId, err)
		return nil, err
	} else {
		fmt.Println(toString(resp.Task))
	}
	return resp.Task, nil
}

func listTasks(resp *ListTaskResponse) {
	for _, task := range resp.Tasks {
		fmt.Println(strconv.Itoa(int(task.TaskId)) + ". " + task.Title)
	}
}

func toString(t *Task) string {
	return "Title: " + t.Title + "\n" +
		"Description: " + t.Description + "\n" +
		"Due date: " + t.DueDate + "\n" +
		"Status: " + t.Status.String() + "\n"
}

func (client TaskAppClient) Close() {
	_ = client.conn.Close()
}
