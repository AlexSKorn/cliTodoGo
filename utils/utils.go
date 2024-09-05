package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

func ReadCsvFile(filename string) ([][]string, error) {

	filePath := getFilePath(filename)

	f, err := loadFile(filePath)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile(f)

	return records, nil
}

func WriteAllToCsv(filename string, rows [][]string) error {
	filePath := getFilePath(filename)

	file, err := loadFile(filePath)
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)

	defer closeFile(file)

	if err := writer.WriteAll(rows); err != nil {
		return fmt.Errorf("error occured writing to file", err)
	}

	return nil
}

func loadFile(filepath string) (*os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

	// Exclusive lock obtained on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}
	return f, nil
}

func closeFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}

func getFilePath(filename string) string {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("error getting current directory")
	}

	filePath := filepath.Join(currentDir, ".", filename)

	return filePath
}
