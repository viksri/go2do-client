package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"todo-client/service"
)

var inputReader = bufio.NewReader(os.Stdin)

type Cli struct {
	client service.TaskAppClient
}

func NewCli(client service.TaskAppClient) Cli {
	return Cli{client: client}
}

func (cli *Cli) askForOperation() int {
	fmt.Println("Choose one of below options: ")
	fmt.Println("1. List existing tasks")
	fmt.Println("2. Create new task")
	fmt.Println("3. Get details for a task")
	fmt.Println("4. Update a task")
	fmt.Println("5. Exit")
	fmt.Printf("Enter your choice(1-5): ")
	input, err := strconv.Atoi(cli.takeInput())
	if err != nil {
		return -1
	}
	return input
}

func (cli *Cli) createTaskPrompt() {
	fmt.Printf("Enter task title: ")
	title := cli.takeInput()
	fmt.Printf("Enter task description (Press ENTER to skip): ")
	desc := cli.takeInput()
	fmt.Printf("Enter task due date in format YYYY-MM-DD HH:mm (Press ENTER to skip): ")
	dueDate := cli.takeInput()

	cli.client.Create(title, desc, dueDate)
}

func (cli *Cli) listTasksPrompt() {
	cli.client.List()
}

func (cli *Cli) detailsTaskPrompt() {
	fmt.Println("Existing tasks are: ")
	cli.listTasksPrompt()
	fmt.Printf("Enter task id for details: ")
	input := cli.takeInput()
	taskId, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Invalid task id: %d\n", input)
	}
	cli.client.Details(int32(taskId))
}

//func (cli *Cli) updateTaskPrompt() {
//	fmt.Println("Existing tasks are: ")
//	cli.listTasksPrompt()
//	fmt.Printf("Enter task id to update: ")
//	input := cli.takeInput()
//	taskId, err := strconv.Atoi(input)
//	if err != nil {
//		fmt.Printf("Invalid task id: %d\n", input)
//	}
//
//	fmt.Printf("Enter task title (Press ENTER to skip): ")
//	title := cli.takeInput()
//	fmt.Printf("Enter task description (Press ENTER to skip): ")
//	desc := cli.takeInput()
//	fmt.Printf("Enter task due date in format YYYY-MM-DD HH:mm (Press ENTER to skip): ")
//	dueDate := cli.takeInput()
//	fmt.Printf("Enter task status (Press ENTER to skip): ")
//	status := cli.takeInput()
//	//cli.client.UpdateTask(int32(taskId), &title, &desc, &dueDate)
//}

func (cli *Cli) takeInput() string {
	input, _ := inputReader.ReadString('\n')
	return strings.TrimSuffix(input, "\n")
}

func (cli *Cli) Start() {
	for {
		choice := cli.askForOperation()
		if choice == 2 {
			cli.createTaskPrompt()
		} else if choice == 1 {
			cli.listTasksPrompt()
		} else if choice == 4 {
			fmt.Println("Goodbye!!!")
			break
		} else if choice == 3 {
			cli.detailsTaskPrompt()
		} else {
			fmt.Printf("Invalid choice: %d\n", choice)
			continue
		}
		fmt.Println("-----------------------")
		time.Sleep(2 * time.Second)
	}
}
