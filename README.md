# go-cli-todo
# Todo CLI App

This is a project for a task list (to-do) application developed in Go. The app runs from the terminal and allows you to manage tasks using different commands. Each task is uniquely identified by a UUID.

## Features

- Add new tasks.
- List all tasks.
- Complete or delete existing tasks.
- Store tasks.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/JohnPosada/go-cli-todo.git
   cd todo-cli-app

2. Install dependencies:

   ```bash  
   go mod download

3. Run the application:

   ```bash
   go run main.go
   ```

## Usage

1. Add a new task:

   ```bash
   go-cli-todo add "My new task"
   ```

2. List all tasks:

   ```bash
   go-cli-todo list
   ```

3. Mark a task as completed:

   ```bash
   go-cli-todo done [UUID]
   ```

4. Delete a task:

   ```bash  
   go-cli-todo delete [UUID]
   ```

## Storage

The application uses a text file to store the tasks. The file path is `./task.json`.


