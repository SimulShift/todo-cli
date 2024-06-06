package cmd

import (
	"github.com/segmentio/ksuid"
	"github.com/spf13/cobra"
	"todo-cli/store"
)

var addCmd = &cobra.Command{
	Use:   "a",
	Short: "Add a new todo",
	Long:  `Add a new todo to the list`,
	Run: func(cmd *cobra.Command, args []string) {
		// make todo struct
		todo := &store.Todo{
			// generate ksuid
			Id:        ksuid.New(),
			Title:     args[0],
			Completed: false,
		}
		store.MemoryStore.AddTodo(*todo)
		store.MemoryStore.Save()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
