/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"todo-cli/cmd"
	"todo-cli/store"
)

func main() {
	store.MemoryStore.Load()
	cmd.Execute()
}
