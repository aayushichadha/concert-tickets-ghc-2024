package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Global logger instance
var Logger = logrus.New()

// init function to configure the logger
func init() {

	// Set logger configuration
	Logger.SetFormatter(&logrus.JSONFormatter{}) // Use JSON format for logs
	Logger.SetOutput(os.Stdout)                  // Output logs to stdout
	Logger.SetLevel(logrus.InfoLevel)            // Set log level to Info

	// Log that initialization is complete
	Logger.Info("Logger initialized in logging package")
}
