package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	jcnfg "github.com/uber/jaeger-client-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	initJaeger("ticket-registry")

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

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Printf("Db connection established")
	return db, nil
}

func initJaeger(service string) {
	cfg, err := jcnfg.FromEnv()
	if err != nil {
		log.Fatalf("Error reading Jaeger config from env: %v", err)
	}

	tracer, closer, err := cfg.New(service)
	if err != nil {
		log.Fatalf("Error creating Jaeger tracer: %v", err)
	}
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
}
