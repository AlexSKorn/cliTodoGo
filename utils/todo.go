package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

func AddTask(description string) {

	records, err := ReadCsvFile("test.csv")
	if err != nil {
		log.Fatal("Error reading file")
	}

	row := []string{
		strconv.Itoa(len(records)),
		description,
		time.Now().Format("2006-01-02"),
		strconv.FormatBool(false),
	}

	records = append(records, row)

	writeErr := WriteAllToCsv("test.csv", records)
	if writeErr != nil {
		log.Fatal("Erorr writing to file in add task", writeErr)
	}
}

func DeleteTask(taskId string) {
	records, err := ReadCsvFile("test.csv")
	if err != nil {
		log.Fatal("Error reading file")
	}

	if err != nil {
		log.Fatal("Error parsing id in delete task", err)
	}

	var updatedRecords [][]string
	for _, record := range records {
		if record[0] != taskId {
			updatedRecords = append(updatedRecords, record)
		}

		if err := WriteAllToCsv("test.csv", updatedRecords); err != nil {
			log.Fatal("Error writing to file in delete task", err)
		}

	}
}

func ListTasks() {
	records, err := ReadCsvFile("test.csv")

	if err != nil {
		fmt.Println("Error reading file")
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
}
