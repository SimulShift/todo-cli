package store

import (
	"github.com/segmentio/ksuid"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

// TodoStore is a struct that holds the store for the list
type TodoStore struct {
	store map[ksuid.KSUID]Todo
}

type Todo struct {
	Id        ksuid.KSUID `yaml:"id"`
	Title     string      `yaml:"title"`
	Completed bool        `yaml:"completed"`
	message   string      `yaml:"message"`
}

var MemoryStore = &TodoStore{
	store: make(map[ksuid.KSUID]Todo),
}

// AddTodo adds a new todo to the store
func (ts *TodoStore) AddTodo(t Todo) {
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
	filePath := filepath.Join(homeDir, ".todo-store.yaml")

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
	filePath := filepath.Join(homeDir, ".todo-store.yaml")

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

// ListTodos lists all the todos in the store
func (ts *TodoStore) ListTodos() {
	for _, todo := range ts.store {
		if todo.Completed {
			println("✅", todo.Title)
		} else {
			println("❌", todo.Title)
		}
	}
}
