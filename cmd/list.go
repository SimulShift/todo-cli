package cmd

import (
	"github.com/spf13/cobra"
	"todo/store"
)

// this command lists the todos

var listCmd = &cobra.Command{
	Use:   "l",
	Short: "List all todos",
	Long:  `List all todos in the list`,
	Run: func(cmd *cobra.Command, args []string) {
		store.MemoryStore.PrintTodos()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
