package main

import (
	"book-tickets/config"
	"book-tickets/gateways"
	"book-tickets/handler"
	"book-tickets/routes"
	"book-tickets/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	jcnfg "github.com/uber/jaeger-client-go/config"
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

	initJaeger("book-tickets")

	ticketRegistryGateway, err := gateways.NewTicketRegistryGateway()
	if err != nil {
		log.Fatalf("Could not create TicketRegistryGateway: %v", err)
	}
	paymentGateway, err := gateways.NewPaymentGateway()
	if err != nil {
		log.Fatalf("Could not create PaymentGateway: %v", err)
	}

	bookingService := &service.BookingService{
		CatalogGateway: ticketRegistryGateway, // Use the TicketRegistryGateway here
		PaymentGateway: paymentGateway,
	}

	bookingHandler := &handler.BookingHandler{
		BookingService: bookingService,
		DB:             db,
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

func initJaeger(service string) {
	// cfg, err := jcnfg.FromEnv()
	// cfg := // Configure Jaeger
	cfg := jcnfg.Configuration{
		ServiceName: "book-tickets", // Change this for each service
		Sampler: &jcnfg.SamplerConfig{
			Type:  "const",
			Param: 1, // Always sample
		},
		Reporter: &jcnfg.ReporterConfig{
			LogSpans: true,
			// Optionally configure agent host and port here
			// LocalAgentHostPort: "jaeger-agent:6831", // Use appropriate host and port
		},
	}
	// if err != nil {
	// 	log.Fatalf("Error reading Jaeger config from env: %v", err)
	// }

	// tracer, closer, err := cfg.New(service)
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Fatalf("Error creating Jaeger tracer: %v", err)
	}
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
}
