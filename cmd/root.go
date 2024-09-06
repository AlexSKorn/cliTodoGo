/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cliTodo/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cliTodo",
	Short: "This is a todo list application to learn how to use cobra",
	Long: `There are 4 commands in this application:
	1. add <description>: To add a new task to the list
	2. list: To list all the tasks in the list
	3. complete <taskid>: complete a task in the list
	4. delete <taskid>: delete a task in the list`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the tasks in the list",
	Long:  `List all the tasks in the list`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.ListTasks()
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the list",
	Long:  `Enter the task description to add to the list`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding task with description: ", args[0])
		utils.AddTask(args[0])
	},
}

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task in the list",
	Long:  `Enter the task id to complete the task`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Completing: ", args[0])
		utils.CompleteTask(args[0])
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task in the list",
	Long:  `Enter the task id to delete the task`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting: ", args[0])
		utils.DeleteTask(args[0])
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cliTodo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(deleteCmd)
}
