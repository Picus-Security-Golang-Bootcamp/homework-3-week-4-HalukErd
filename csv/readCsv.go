package csv

import (
	"encoding/csv"
	"github.com/HalukErd/Week4Assignment/models"
	"os"
)

// ReadCsvToBookAndAuthor read csv to a BookCsvLines struct
func ReadCsvToBookAndAuthor(fName string) (models.BookCsvLines, error) {
	f, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var bookCsvLines models.BookCsvLines

	for _, line := range records[1:] {
		bookCsvLines = append(bookCsvLines, models.BookCsvLine{
			Title:     line[0],
			Author:    line[1],
			Genre:     line[2],
			Height:    line[3],
			Publisher: line[4],
		})
	}
	return bookCsvLines, nil
}
