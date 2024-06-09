package cmd

import (
	"github.com/segmentio/ksuid"
	"github.com/spf13/cobra"
	"todo-cli/store"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a todo as done",
	Long:  `Mark a todo as done`,
	Run: func(cmd *cobra.Command, args []string) {
		store.MemoryStore.Load()
		// check arg[0] is a valid ksuid
		id, err := ksuid.Parse(args[0])
		if err != nil {
			panic(err)
		}
		store.MemoryStore.Done(id)
		store.MemoryStore.Save()
	},
}
