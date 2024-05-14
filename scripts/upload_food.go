package scripts

import (
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/newtri-science/synonym-tool/model"
)


func GenerateFoodEntries(file *multipart.FileHeader) ([]model.Food, error) {
	// test for .csv
    if !strings.HasSuffix(file.Filename, ".csv") {
		return nil, fmt.Errorf("file is no .csv file")
    }

	// open file
    src, err := file.Open()
    if err != nil {
        return nil, fmt.Errorf("error while opening the csv file: %s", err)
    }
    defer src.Close()

    // read file content
    content, err := io.ReadAll(src)
    if err != nil {
        return nil, fmt.Errorf("error while reading the csv stream: %s", err)
    }

    // read csv
    csvReader := csv.NewReader(strings.NewReader(string(content)))
	csvReader.Comma = ';'
    records, err := csvReader.ReadAll()
    if err != nil {
        return nil, fmt.Errorf("error while reading the csv file, please use SEMICOLON as separator: %s", err)
    }

	fmt.Println(records)

    // Erfolgreiche Verarbeitung der CSV-Datei
    // records enthält nun den Inhalt der CSV-Datei als 2D-Slice

	return nil, nil
}

// Überprüfe ob .csv und ob die richtigen Columns

// Bring alles in Form (sonderzeichen müssen raus)

// Schreibe in DB (KEINE doppelten Namen und KEINE doppelten IDs)

// Bei doppelten IDs oder Namen: lösche alte Einträge und schreibe neue
