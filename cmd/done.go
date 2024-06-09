package cmd

import (
	"fmt"
	"github.com/segmentio/ksuid"
	"github.com/spf13/cobra"
	"todo/store"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a todo as done",
	Long:  `Mark a todo as done`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a todo ID")
			return
		}
		store.MemoryStore.Load()
		// check arg[0] is a valid ksuid
		id, err := ksuid.Parse(args[0])
		if err != nil {
			panic(err)
		}
		found := store.MemoryStore.DoneById(id)
		if found {
			store.MemoryStore.Save()
			return
		}
		// Try to find the todo by title
		store.MemoryStore.DoneByTitle(args[0])
		store.MemoryStore.Save()
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// Load current todos from the memory store
		store.MemoryStore.Load()
		todos := store.MemoryStore.GetAllTodos()
		var completions []string
		// Filter todos that are not done yet and collect their IDs
		for _, todo := range todos {
			if !todo.Done {
				completions = append(completions, todo.Id.String())
			}
		}
		return completions, cobra.ShellCompDirectiveNoFileComp
	},
}

// Assuming 'store' has a method to list all todo items
func init() {
	rootCmd.AddCommand(doneCmd)
}
