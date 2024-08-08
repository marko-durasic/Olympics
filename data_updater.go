package main

import (
	"context"
	"log"

	"github.com/marko-durasic/olympics/ent"
)

type DataUpdater struct {
	client *ent.Client
}

func NewDataUpdater(client *ent.Client) *DataUpdater {
	return &DataUpdater{client: client}
}

func (du *DataUpdater) UpdateTeamSports(ctx context.Context) {
	log.Println("Updating team sports data...")

	// Clear existing data
	_, err := du.client.TeamSport.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("failed clearing team sports data: %v", err)
	}

	// Read data from CSV
	data, err := ReadCSV("sample_data/sample_data.csv")
	if err != nil {
		log.Fatalf("failed reading CSV data: %v", err)
	}

	// Insert data from CSV
	for _, sport := range data {
		_, err = du.client.TeamSport.Create().
			SetCountry(sport.Country).
			SetSport(sport.Sport).
			SetGold(sport.Gold).
			SetSilver(sport.Silver).
			SetBronze(sport.Bronze).
			SetPoints(sport.Points).
			SetTotalScore(sport.TotalScore).
			SetPopulation(sport.Population).
			SetPerCapita(sport.PerCapita).Save(ctx)
		if err != nil {
			log.Fatalf("failed inserting team sports data: %v", err)
		}
	}

	log.Println("Team sports data updated.")
}

func (du *DataUpdater) UpdateIndividualSports(ctx context.Context) {
	log.Println("Updating individual sports data...")

	// Clear existing data
	_, err := du.client.IndividualSport.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("failed clearing individual sports data: %v", err)
	}

	// Read data from CSV
	data, err := ReadCSV("sample_data/sample_data.csv")
	if err != nil {
		log.Fatalf("failed reading CSV data: %v", err)
	}

	// Insert data from CSV
	for _, sport := range data {
		if sport.Sport == "Swimming" { // Adjust condition as needed
			_, err = du.client.IndividualSport.Create().
				SetCountry(sport.Country).
				SetSport(sport.Sport).
				SetGold(sport.Gold).
				SetSilver(sport.Silver).
				SetBronze(sport.Bronze).
				SetPoints(sport.Points).
				SetTotalScore(sport.TotalScore).
				SetPopulation(sport.Population).
				SetPerCapita(sport.PerCapita).Save(ctx)
			if err != nil {
				log.Fatalf("failed inserting individual sports data: %v", err)
			}
		}
	}

	log.Println("Individual sports data updated.")
}

func (du *DataUpdater) UpdateCombinedSports(ctx context.Context) {
	log.Println("Updating combined sports data...")

	// Clear existing data
	_, err := du.client.CombinedSport.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("failed clearing combined sports data: %v", err)
	}

	// Read data from CSV
	data, err := ReadCSV("sample_data/sample_data.csv")
	if err != nil {
		log.Fatalf("failed reading CSV data: %v", err)
	}

	// Insert data from CSV
	for _, sport := range data {
		if sport.Sport == "Athletics" { // Adjust condition as needed
			_, err = du.client.CombinedSport.Create().
				SetCountry(sport.Country).
				SetSport(sport.Sport).
				SetGold(sport.Gold).
				SetSilver(sport.Silver).
				SetBronze(sport.Bronze).
				SetPoints(sport.Points).
				SetTotalScore(sport.TotalScore).
				SetPopulation(sport.Population).
				SetPerCapita(sport.PerCapita).Save(ctx)
			if err != nil {
				log.Fatalf("failed inserting combined sports data: %v", err)
			}
		}
	}

	log.Println("Combined sports data updated.")
}

func (du *DataUpdater) UpdateAll(ctx context.Context) {
	go du.UpdateTeamSports(ctx)
	go du.UpdateIndividualSports(ctx)
	go du.UpdateCombinedSports(ctx)
}
