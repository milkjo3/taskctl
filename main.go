//===============================================
// taskctl v1.0
// Author: @milkjo3
// Date: 2025-06-06
// Description: A simple CLI Task Management Tool
// License: MIT
//===============================================

package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

const logo = `
████████╗ █████╗ ███████╗██╗  ██╗ ██████╗████████╗██╗     
╚══██╔══╝██╔══██╗██╔════╝██║ ██╔╝██╔════╝╚══██╔══╝██║     
   ██║   ███████║███████╗█████╔╝ ██║        ██║   ██║     
   ██║   ██╔══██║╚════██║██╔═██╗ ██║        ██║   ██║     
   ██║   ██║  ██║███████║██║  ██╗╚██████╗   ██║   ███████╗
   ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚══════╝
           A simple CLI Task Management Tool (taskctl)` + "\n"

// Colors
const (
	green = "\033[32m"
	red   = "\033[31m"
	reset = "\033[0m"
)

type Task struct {
	Name     string `json:"name"`
	Done     bool   `json:"done"`
	Priority string `json:"priority"`
	DueDate  string `json:"dueDate"`
}

// TaskCtlr struct. This is the main struct that will be used to store the tasks.
// CRUD operations will be performed on this struct.
type taskCtlr struct {
	Tasks    []Task
	Filename string
}

// createTask function. This function will create a new task.
func createTask(tm *taskCtlr) {
	fmt.Print("Enter the task name: ")
	reader := bufio.NewReader(os.Stdin)
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)

	// Ask user for priority and due date
	taskPriority := getValidPriority()
	taskDueDate := getValidDueDate()
	task := Task{Name: taskName, Done: false, Priority: taskPriority, DueDate: taskDueDate}

	fmt.Println("Task added!")
	tm.Tasks = append(tm.Tasks, task)
	saveTasks(tm)
}

// readTask function. This function will read a task.
func readTask(tm *taskCtlr, index int) {
	// Check if index is valid
	if index < 0 || index >= len(tm.Tasks) {
		fmt.Println("Invalid task index")
		return
	}

	// Print task details
	status := "Not Done"
	color := red
	if tm.Tasks[index].Done {
		status = "Done"
		color = green
	}

	due := tm.Tasks[index].DueDate
	if due == "" {
		due = "(none)"
	}

	fmt.Printf("\n%-4s | %-20s | %-8s | %-8s | %-10s\n", "ID", "Name", "Status", "Priority", "Due Date")
	fmt.Println("-----+----------------------+----------+----------+------------")
	fmt.Printf("%-4d | %-20s | %s%-8s%s | %-8s | %-10s\n",
		index,
		tm.Tasks[index].Name,
		color, status, reset,
		strings.Title(tm.Tasks[index].Priority),
		due)
}

// toggleTaskDone function. This function will toggle the done status of a task.
func toggleTaskDone(tm *taskCtlr, index int) {
	fmt.Println("Task updated!")
	tm.Tasks[index].Done = !tm.Tasks[index].Done
	saveTasks(tm)
}

// deleteTask function. This function will delete a task.
func deleteTask(tm *taskCtlr, index int) error {
	// Check if index is valid
	if index < 0 || index >= len(tm.Tasks) {
		return errors.New("invalid task index")
	}

	// Delete task. Append the task before the index and the task after the index.
	tm.Tasks = append(tm.Tasks[:index], tm.Tasks[index+1:]...)
	fmt.Println("Task deleted!")
	saveTasks(tm)
	return nil
}

// loadTasks function. This function will load the tasks from the JSON file.
func loadTasks(tm *taskCtlr) error {
	jsonData, err := os.ReadFile(tm.Filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, &tm.Tasks)
	if err != nil {
		return err
	}
	return nil
}

// saveTasks function. This function will save the tasks to the JSON file.
func saveTasks(tm *taskCtlr) error {
	jsonData, err := json.MarshalIndent(tm.Tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(tm.Filename, jsonData, 0644)
}

// clearTasks function. This function will clear all tasks.
func clearTasks(tm *taskCtlr) {
	tm.Tasks = []Task{}
	saveTasks(tm)
	fmt.Println("Tasks cleared!")
}

// viewAllTasks function. This function will view all tasks.
func viewAllTasks(tm *taskCtlr) {
	// Check if there are any tasks
	if len(tm.Tasks) == 0 {
		fmt.Println("No tasks found.")
	} else {
		fmt.Printf("\n%-4s | %-20s | %-8s | %-8s | %-10s\n", "ID", "Name", "Status", "Priority", "Due Date")
		fmt.Println("-----+----------------------+----------+----------+------------")

		for index, task := range tm.Tasks {
			status := "Not Done"
			color := red
			if task.Done {
				status = "Done"
				color = green
			}
			due := task.DueDate
			if due == "" {
				due = "(none)"
			}
			fmt.Printf("%-4d | %-20s | %s%-8s%s | %-8s | %-10s\n",
				index,
				task.Name,
				color, status, reset,
				strings.Title(task.Priority),
				due)
		}
	}
}

// getValidIndex function. This function will get a valid index from the user.
// Used for read, update, and delete tasks. Cannot be used when no tasks are found.
func getValidIndex(tm *taskCtlr) int {
	// Check if there are any tasks
	if len(tm.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return -1
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Enter the task index (0-%d): ", len(tm.Tasks)-1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		index, err := strconv.Atoi(input)
		if err != nil || index < 0 || index >= len(tm.Tasks) {
			fmt.Println("Invalid task index. Please try again.")
			continue
		}
		return index
	}
}

// getValidPriority function. This function will get a priority level from the user for a task.
func getValidPriority() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter the task priority (low, medium, high): ")
		input, _ := reader.ReadString('\n')
		priority := strings.ToLower(strings.TrimSpace(input))
		if priority == "low" || priority == "medium" || priority == "high" {
			return priority
		}
		fmt.Println("Invalid priority. Please try again.")
	}
}

// getValidDueDate function. This function will get a due date from the user for a task.
func getValidDueDate() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter the task due date (YYYY-MM-DD), or press Enter to skip: ")
		input, _ := reader.ReadString('\n')
		dueDate := strings.TrimSpace(input)

		// If the user presses Enter, return an empty string. No due date for the task.
		if dueDate == "" {
			return ""
		}

		_, err := time.Parse("2006-01-02", dueDate)
		if err != nil {
			fmt.Println("Invalid date format. Please try again.")
			continue
		}
		return dueDate
	}
}

// viewByPriority function. This function will view all tasks sorted by priority.
func viewByPriority(tm *taskCtlr) {
	if len(tm.Tasks) == 0 {
		fmt.Println("No tasks found.")
	} else {
		var priorityWeight = map[string]int{
			"high":   1,
			"medium": 2,
			"low":    3,
		}

		tasks := make([]Task, len(tm.Tasks))
		copy(tasks, tm.Tasks)

		sort.Slice(tasks, func(i, j int) bool {
			return priorityWeight[tasks[i].Priority] < priorityWeight[tasks[j].Priority]
		})

		fmt.Printf("\n%-4s | %-20s | %-8s | %-8s | %-10s\n", "ID", "Name", "Status", "Priority", "Due Date")
		fmt.Println("-----+----------------------+----------+----------+------------")

		for index, task := range tasks {
			status := "Not Done"
			color := red
			if task.Done {
				status = "Done"
				color = green
			}
			due := task.DueDate
			if due == "" {
				due = "(none)"
			}
			fmt.Printf("%-4d | %-20s | %s%-8s%s | %-8s | %-10s\n",
				index,
				task.Name,
				color, status, reset,
				strings.Title(task.Priority),
				due)
		}
	}
}

// viewByStatus function. This function will view all tasks sorted by status.
func viewByStatus(tm *taskCtlr) {
	if len(tm.Tasks) == 0 {
		fmt.Println("No tasks found.")
	} else {
		var priorityWeight = map[string]int{
			"true":  2,
			"false": 1,
		}

		tasks := make([]Task, len(tm.Tasks))
		copy(tasks, tm.Tasks)

		sort.Slice(tasks, func(i, j int) bool {
			return priorityWeight[strconv.FormatBool(tasks[i].Done)] < priorityWeight[strconv.FormatBool(tasks[j].Done)]
		})

		fmt.Printf("\n%-4s | %-20s | %-8s | %-8s | %-10s\n", "ID", "Name", "Status", "Priority", "Due Date")
		fmt.Println("-----+----------------------+----------+----------+------------")

		for index, task := range tasks {
			status := "Not Done"
			color := red
			if task.Done {
				status = "Done"
				color = green
			}
			due := task.DueDate
			if due == "" {
				due = "(none)"
			}
			fmt.Printf("%-4d | %-20s | %s%-8s%s | %-8s | %-10s\n",
				index,
				task.Name,
				color, status, reset,
				strings.Title(task.Priority),
				due)
		}
	}
}

// viewByDueDate function. This function will view all tasks sorted by due date.
func viewByDueDate(tm *taskCtlr) {
	if len(tm.Tasks) == 0 {
		fmt.Println("No tasks found.")
	} else {
		tasks := make([]Task, len(tm.Tasks))
		copy(tasks, tm.Tasks)

		layout := "2006-01-02"

		sort.Slice(tasks, func(i, j int) bool {
			dateI, errI := time.Parse(layout, tasks[i].DueDate)
			dateJ, errJ := time.Parse(layout, tasks[j].DueDate)

			if errI != nil {
				return false
			}

			if errJ != nil {
				return true
			}

			return dateI.Before(dateJ)
		})

		fmt.Printf("\n%-4s | %-20s | %-8s | %-8s | %-10s\n", "ID", "Name", "Status", "Priority", "Due Date")
		fmt.Println("-----+----------------------+----------+----------+------------")

		for index, task := range tasks {
			status := "Not Done"
			color := red
			if task.Done {
				status = "Done"
				color = green
			}
			due := task.DueDate
			if due == "" {
				due = "(none)"
			}
			fmt.Printf("%-4d | %-20s | %s%-8s%s | %-8s | %-10s\n",
				index,
				task.Name,
				color, status, reset,
				strings.Title(task.Priority),
				due)
		}
	}
}

// pause function. This function will pause the program and wait for the user to press Enter.
func pause() {
	// Ensure output buffer is flushed
	fmt.Print("\nPress Enter to continue...")
	os.Stdout.Sync()

	reader := bufio.NewReader(os.Stdin)

	// Wait for the Enter key
	_, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}

	clearScreen()
}

func clearScreen() {
	// Add a small delay before clearing to ensure all output is visible
	time.Sleep(50 * time.Millisecond)

	// clear screen using ANSI escape codes
	fmt.Print("\033[H\033[2J")
	os.Stdout.Sync()
}

// promptMainMenu function. This function will prompt the user to select an option from the main menu.
func promptMainMenu() int {

	printLogo()

	// prompt the user to select an option from the main menu
	prompt := promptui.Select{
		Label: "Welcome to the Task Manager",
		Items: []string{
			"Create | Read | Update | Delete Tasks",
			"View Tasks",
			"Clear Tasks",
			"Exit",
		},
		Size: 4,
	}

	index, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1
	}

	return index + 1
}

// promptTaskMenu function. This function will prompt the user to select an option from the task menu.
func promptTaskMenu() int {

	printLogo()
	// prompt the user to select an option from the task menu
	prompt := promptui.Select{
		Label: "Task Management Menu",
		Items: []string{
			"Create Task",
			"Read Task",
			"Update Task",
			"Delete Task",
			"Go back to main menu",
		},
		Size: 5,
	}

	index, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1
	}
	return index + 1
}

// promptViewMenu function. This function will prompt the user to select an option from the view menu.
func promptViewMenu() int {

	printLogo()

	// prompt the user to select an option from the view menu
	prompt := promptui.Select{
		Label: "View Tasks Menu",
		Items: []string{
			"View All",
			"View by Priority",
			"View by Status",
			"View by Due Date",
			"Go Back to main menu",
		},
		Size: 5,
	}

	index, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1
	}
	return index + 1
}

// printLogo function. This prints the most beautiful logo ever made. No question asked.
func printLogo() {
	clearScreen()
	fmt.Print(logo)
}

// main function. This function will run the program.
func main() {
	var choice int
	var index int
	tm := taskCtlr{Filename: "tasks.json"}
	loadTasks(&tm)

	// main loop
	for {
		choice = promptMainMenu()
		switch choice {
		case 1:
		taskMenuLoop:
			for {
				choice = promptTaskMenu()
				switch choice {
				case 1:
					printLogo()
					createTask(&tm)
					pause()
				case 2:
					clearScreen()
					index = getValidIndex(&tm)
					if index == -1 {
						pause()
						break
					}
					readTask(&tm, index)
					pause()
				case 3:
					clearScreen()
					index = getValidIndex(&tm)
					if index == -1 {
						pause()
						break
					}
					toggleTaskDone(&tm, index)
					pause()
				case 4:
					clearScreen()
					index = getValidIndex(&tm)
					if index == -1 {
						pause()
						break
					}
					deleteTask(&tm, index)
					pause()
				case 5:
					break taskMenuLoop
				default:
					fmt.Println("Invalid choice")
					pause()
				}

			}
		case 2:
		viewMenuLoop:
			for {
				choice = promptViewMenu()
				switch choice {
				case 1:
					printLogo()
					viewAllTasks(&tm)
					pause()
				case 2:
					printLogo()
					viewByPriority(&tm)
					pause()
				case 3:
					printLogo()
					viewByStatus(&tm)
					pause()
				case 4:
					printLogo()
					viewByDueDate(&tm)
					pause()
				case 5:
					break viewMenuLoop
				default:
					fmt.Println("Invalid choice")
					pause()
				}

			}
		case 3:
			clearScreen()
			clearTasks(&tm)
			pause()
		case 4:
			clearScreen()
			fmt.Println("Exiting... Thanks for using taskctl!")
			return
		default:
			fmt.Println("Invalid choice")
			pause()
		}
	}
}
