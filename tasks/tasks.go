package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
}

// LoadTask reads the given file and unmarshals the content into a slice of Task.
// If the file is empty, it returns an empty slice. If there is an error while reading
// the file or unmarshaling the content, it panics.
func LoadTask(file *os.File) []Task {
	// Initialize an empty slice of Task
	tasks := make([]Task, 0)

	// Get file statistics
	fileStat, err := file.Stat()
	if err != nil {
		panic(err) // Panic if there is an error getting the file stats
	}

	// Return empty tasks slice if the file is empty
	if fileStat.Size() == 0 {
		return tasks
	}

	// Read all content from the file
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err) // Panic if there is an error reading the file
	}

	// Unmarshal the JSON content into the tasks slice
	if err := json.Unmarshal(fileBytes, &tasks); err != nil {
		panic(err) // Panic if there is an error unmarshaling the JSON
	}

	// Return the list of tasks
	return tasks
}

// List prints each task in the given list in a human-readable format. Completed tasks will have a
// "✓" in the first column, while incomplete tasks will have a blank in that column.
func List(tasks []Task) {
	for _, task := range tasks {
		completed := " "
		if task.Completed {
			completed = "✓"
		}
		fmt.Printf("[%s] -- %s -- %s\n", completed, task.ID, task.Name)
	}
}

// Add creates a new Task with the given name and adds it to the given list of tasks. The new task is
// not completed. The given list of tasks is modified and the modified list is returned.
func Add(tasks []Task, name string) []Task {
	newTask := Task{
		ID:        uuid.New(),
		Name:      name,
		Completed: false,
	}
	tasks = append(tasks, newTask)

	return tasks
}

// Save takes a file and a list of tasks and overwrites the file with the tasks as JSON, flushing the
// writer at the end. If the write fails, it panics.
func Save(file *os.File, tasks []Task) {
	jsonBytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	// Seek the beginning of the file, truncate it, and write the new JSON data
	file.Seek(0, 0)
	file.Truncate(0)

	// Use a buffered writer to write the JSON data
	writer := bufio.NewWriter(file)
	_, err = writer.Write(jsonBytes)
	if err != nil {
		// If the write fails, panic
		panic(err)
	}

	// Flush the buffered writer to make sure everything is written to disk
	writer.Flush()
}

// Done toggles the completion status of a task with the given id in the list of tasks.
func Done(tasks []Task, id uuid.UUID) []Task {
	for i, task := range tasks {
		if task.ID == id {
			// Toggle the completion status of the task
			tasks[i].Completed = !tasks[i].Completed
			break
		}
	}
	return tasks
}

// Delete removes a task with the given id from the list of tasks.
func Delete(tasks []Task, id uuid.UUID) []Task {
	// Iterate through the tasks to find the task with the given id
	for i, task := range tasks {
		if task.ID == id {
			// Remove the task from the list by slicing it out
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	// Return the updated tasks list
	return tasks
}
