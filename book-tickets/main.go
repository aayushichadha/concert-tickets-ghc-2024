package main

import (
	"book-tickets/config"
	"book-tickets/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"path/filepath"
)

func main() {

	// Load configuration
	configPath, err := filepath.Abs("config/config.json")
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		db, err := setupDB(*cfg)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}
		c.Set("db", db) // Store the DB connection in the context
		c.Next()
	})
	routes.SetupRoutes(router)
	router.Run(":8083") // listen and serve on 0.0.0.0:8083

}

func setupDB(config config.Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host, config.Database.Port, config.Database.User,
		config.Database.Password, config.Database.DBName, config.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Printf("Db connection established")
	return db, nil
}
