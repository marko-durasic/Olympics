// csv_reader.go
package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type SportData struct {
	Country    string
	Sport      string
	Gold       int
	Silver     int
	Bronze     int
	Points     int
	TotalScore int
	Population int
	PerCapita  float64
}

func ReadCSV(filePath string) ([]SportData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []SportData
	for _, record := range records[1:] { // Skip header row
		gold, _ := strconv.Atoi(record[2])
		silver, _ := strconv.Atoi(record[3])
		bronze, _ := strconv.Atoi(record[4])
		points, _ := strconv.Atoi(record[5])
		totalScore, _ := strconv.Atoi(record[6])
		population, _ := strconv.Atoi(record[7])
		perCapita, _ := strconv.ParseFloat(record[8], 64)

		data = append(data, SportData{
			Country:    record[0],
			Sport:      record[1],
			Gold:       gold,
			Silver:     silver,
			Bronze:     bronze,
			Points:     points,
			TotalScore: totalScore,
			Population: population,
			PerCapita:  perCapita,
		})
	}

	return data, nil
}
