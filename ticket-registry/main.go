package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"path/filepath"
	"ticket-registry/config"
	"ticket-registry/routes"
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
	router.Run(":8082") // listen and serve on 0.0.0.0:8082

}

func setupDB(config config.Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host, config.Database.Port, config.Database.User,
		config.Database.Password, config.Database.DBName, config.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	log.Printf("Db connection established")
	return db, nil
}
