package cmd

import (
	"cliTodo/utils"
	"os"
	"testing"
)

func TestDeleteTask(t *testing.T) {
	tempFile,                                                                                                                                                                                                                                                                                                                                                                           err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	defer os.Remove(tempFile.Name())

	testData := [][]string{
		{"1", "Task 1", "2024-07-27", "false"},
		{"2", "Task 2", "2024-07-28", "false"},
	}

	if err := utils.WriteAllToCsv(tempFile.Name(), testData); err != nil {
		t.Fatalf("Failed to write test data: %v", err)
	}

}
