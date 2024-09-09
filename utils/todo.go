package utils

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

const (
	CsvFileName = "test.csv"
	DateFormat  = "2006-01-02"
)

func AddTask(description string) error {

	records, err := ReadCsvFile(CsvFileName)
	if err != nil {
		return fmt.Errorf("error writing to file in add task: %w", err)
	}

	row := []string{
		strconv.Itoa(len(records)),
		description,
		time.Now().Format(DateFormat),
		strconv.FormatBool(false),
	}

	records = append(records, row)

	writeErr := WriteAllToCsv(CsvFileName, records)
	if writeErr != nil {
		return fmt.Errorf("error writing to file in add task: %w", writeErr)
	}

	return nil
}

func DeleteTask(taskId string) error {
	records, err := ReadCsvFile(CsvFileName)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	var updatedRecords [][]string
	for _, record := range records {
		if record[0] != taskId {
			updatedRecords = append(updatedRecords, record)
		}
	}

	if err := WriteAllToCsv(CsvFileName, updatedRecords); err != nil {
		return fmt.Errorf("error writing to file in delete task: %w", err)
	}

	return nil
}

func ListTasks() error {
	records, err := ReadCsvFile(CsvFileName)

	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	for _, record := range records {
		id := record[0]
		task := record[1]
		created := record[2]
		done := record[3]

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", id, task, created, done)
	}

	w.Flush()
	return nil
}

func CompleteTask(taskId string) error {
	records, err := ReadCsvFile(CsvFileName)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	for _, record := range records {
		if record[0] == taskId {
			record[3] = strconv.FormatBool(true)
		}
	}
	if err := WriteAllToCsv(CsvFileName, records); err != nil {
		return fmt.Errorf("error editing task to complete: %w", err)
	}
	return nil
}
