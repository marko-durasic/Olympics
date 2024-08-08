package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/marko-durasic/olympics/ent"
)

var client *ent.Client
var dataUpdater *DataUpdater

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable not set")
	}

	var err error
	for i := 0; i < 10; i++ {
		drv, err := sql.Open(dialect.Postgres, dbURL)
		if err != nil {
			log.Printf("failed connecting to postgres (attempt %d/10): %v", i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}
		client = ent.NewClient(ent.Driver(drv))
		ctx := context.Background()
		if err = client.Schema.Create(ctx); err != nil {
			log.Printf("failed creating schema resources (attempt %d/10): %v", i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}

	if err != nil {
		log.Fatalf("could not connect to postgres after 10 attempts: %v", err)
	}

	defer client.Close()
	dataUpdater = NewDataUpdater(client)

	// Run data update during startup
	ctx := context.Background()
	dataUpdater.UpdateAll(ctx)

	router := gin.Default()

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.GET("/api/team-sports", getTeamSports)
	router.GET("/api/individual-sports", getIndividualSports)
	router.GET("/api/combined", getCombined)
	router.POST("/api/update-data", updateData)

	router.Run(":8080")
}

func getTeamSports(c *gin.Context) {
	ctx := context.Background()
	sports, err := client.TeamSport.Query().All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sports)
}

func getIndividualSports(c *gin.Context) {
	ctx := context.Background()
	sports, err := client.IndividualSport.Query().All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sports)
}

func getCombined(c *gin.Context) {
	ctx := context.Background()
	sports, err := client.CombinedSport.Query().All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sports)
}

func updateData(c *gin.Context) {
	ctx := context.Background()
	dataUpdater.UpdateAll(ctx)
	c.JSON(http.StatusOK, gin.H{"message": "Data update started"})
}
