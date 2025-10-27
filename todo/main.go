package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"todo/model"
	"todo/storage"
)

var (
	storageType string
	store       storage.Operator
	filePath    string
)

func initStorage() error {
	factory := storage.NewStorageFactory()
	var err error
	filePath = "todo.json"

	if storageType == "yaml" {
		filePath = "todo.yaml"
	}

	store, err = factory.CreateStorage(storageType, filePath)
	return err
}

func main() {

	var rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "A simple todo list",
		Long:  `A simple todo list`,
	}

	var addCmd = &cobra.Command{
		Use:   "add [title]",
		Short: "Add a new todo",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := initStorage(); err != nil {
				fmt.Printf("Error initializing storage: %v\n", err)
				return
			}
			todo := &model.Todo{
				Content: args[0],
				Done:    false,
			}

			if err := store.Add(todo); err != nil {
				fmt.Printf("Error adding todo: %v\n", err)
				return
			}
			fmt.Println("Todo added successfully!")
		},
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all todos",
		Run: func(cmd *cobra.Command, args []string) {
			if err := initStorage(); err != nil {
				fmt.Printf("Error initializing storage: %v\n", err)
				return
			}
			todos, _ := store.List()
			if len(todos) == 0 {
				fmt.Println("No todos found.")
				return
			}
			for _, todo := range todos {
				fmt.Printf("%d %s\n", todo.Id, todo.Content)

			}
		},
	}

	var completeCmd = &cobra.Command{
		Use:   "done [id]",
		Short: "Mark a todo as completed",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := initStorage(); err != nil {
				fmt.Printf("Error initializing storage: %v\n", err)
				return
			}

			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("Invalid ID: %v\n", err)
				return
			}

			if err := store.Done(id); err != nil {
				fmt.Printf("Error updating todo: %v\n", err)
				return
			}
			fmt.Println("Todo marked as completed!")
		},
	}

	rootCmd.PersistentFlags().StringVarP(&storageType, "type", "t", "json", "Storage type (json or yaml)")
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(completeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
