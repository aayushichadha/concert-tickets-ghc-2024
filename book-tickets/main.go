package main

import (
	"book-tickets/config"
	"book-tickets/gateways"
	"book-tickets/handler"
	"book-tickets/routes"
	"book-tickets/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"path/filepath"
)

func main() {
	configPath, err := filepath.Abs("config/config.json")
	if err != nil {
		log.Fatalf("Could not determine config file path: %v", err)
	}
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	db, err := setupDB(*cfg)
	if err != nil {
		log.Fatalf("Could not set up database connection: %v", err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{}) // Use JSON format for structured logs
	logger.SetLevel(logrus.InfoLevel)

	ticketRegistryGateway, err := gateways.NewTicketRegistryGateway(logger)
	if err != nil {
		log.Fatalf("Could not create TicketRegistryGateway: %v", err)
	}
	paymentGateway, err := gateways.NewPaymentGateway(logger)
	if err != nil {
		log.Fatalf("Could not create PaymentGateway: %v", err)
	}

	bookingService := &service.BookingService{
		CatalogGateway: ticketRegistryGateway, // Use the TicketRegistryGateway here
		PaymentGateway: paymentGateway,
		Logger:         logger,
	}

	bookingHandler := &handler.BookingHandler{
		BookingService: bookingService,
		DB:             db,
		Logger:         logger,
	}

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db) // Store the DB connection in the context
		c.Next()
	})

	routes.SetupRoutes(router, bookingHandler)

	router.Run(":8083") // listen and serve on 0.0.0.0:8083
}

// setupDB initializes the PostgreSQL connection using GORM
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
