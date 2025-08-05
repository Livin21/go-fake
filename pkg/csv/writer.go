package csv

import (
    "encoding/csv"
    "os"
)

// WriteCSV writes the provided data to a CSV file at the specified file path.
func WriteCSV(filePath string, data [][]string) error {
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, record := range data {
        if err := writer.Write(record); err != nil {
            return err
        }
    }

    return nil
}