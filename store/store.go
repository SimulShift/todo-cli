package store

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

// TodoStore is a struct that holds the store for the list
type TodoStore struct {
	store map[int]Todo
}

type Todo struct {
	Id      int    `yaml:"id"`
	Title   string `yaml:"title"`
	Done    bool   `yaml:"completed"`
	Message string `yaml:"message"`
}

var MemoryStore = &TodoStore{
	store: make(map[int]Todo),
}

var filename = ".todo-store2.yaml"

func (ts *TodoStore) GetAllTodos() (todos []Todo) {
	// get all the todos from the store
	todoMap := ts.store
	// put the todos in a list
	for _, todo := range todoMap {
		todos = append(todos, todo)
	}
	return todos
}

// GetMaxId returns the maximum id in the store
func (ts *TodoStore) GetMaxId() int {
	maxId := 0
	for _, todo := range ts.store {
		if todo.Id > maxId {
			maxId = todo.Id
		}
	}
	return maxId
}

// AddTodo adds a new todo to the store
func (ts *TodoStore) AddTodo(t Todo) {
	// make sure the title does not exist
	for _, todo := range ts.store {
		if todo.Title == t.Title {
			println("Todo with title already exists")
			return
		}

	}
	maxId := ts.GetMaxId()
	t.Id = maxId + 1
	ts.store[t.Id] = t
}

// Save - Save TodoStore as yaml
func (ts *TodoStore) Save() {
	// Save the store as yaml
	yamlData, err := yaml.Marshal(ts.store)
	if err != nil {
		panic(err)
	}
	// Save the yaml in Home directory

	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err) // Handle the error according to your application's error policy
	}

	// Create the full path to save the file in the home directory
	filePath := filepath.Join(homeDir, filename)

	err = os.WriteFile(filePath, yamlData, 0644)
	if err != nil {
		panic(err)
	}
}

// Load - Load TodoStore from yaml
func (ts *TodoStore) Load() {
	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err) // Handle the error according to your application's error policy
	}

	// Create the full path to save the file in the home directory
	filePath := filepath.Join(homeDir, filename)

	// Load yaml data from file
	yamlData, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Unmarshal yaml data to store
	err = yaml.Unmarshal(yamlData, &ts.store)
	if err != nil {
		panic(err)
	}
}

// PrintTodos lists all the todos in the store
func (ts *TodoStore) PrintTodos() {
	for _, todo := range ts.store {
		if todo.Done {
			println(todo.Id, "✅", todo.Title)
		} else {
			println(todo.Id, "❌", todo.Title)
		}
	}
}

// Done sets a todo as done
func (ts *TodoStore) DoneById(id int) bool {
	todo, ok := ts.store[id]
	if !ok {
		println("Todo not found")
		return false
	}
	todo.Done = true
	ts.store[id] = todo
	return true
}

func (ts *TodoStore) DoneByTitle(title string) {
	for id, todo := range ts.store {
		if todo.Title == title {
			todo.Done = true
			ts.store[id] = todo
		}
	}
	println("Todo not found")
}
