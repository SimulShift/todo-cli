/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"todo/cmd"
	"todo/store"
)

func main() {
	store.MemoryStore.Load()
	cmd.Execute()
}
