package functions

import (
	"github.com/sirupsen/logrus"
	"os"
)

func SetupLogger() {
	// Create a new logger instance
	logger := logrus.New()

	// Set log level (debug, info, warn, error, fatal, panic)
	logger.SetLevel(logrus.DebugLevel)

	// You can also set a custom log format if needed
	// For example, to include timestamp and log level in each log entry:
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	// Output logs to the console
	logger.SetOutput(os.Stdout)

	// You can also log to a file or any other custom output by setting it here
	// For example:
	// file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	//     logger.SetOutput(file)
	// } else {
	//     logger.Info("Failed to log to file, using default stderr")
	// }

	// Save the logger for use throughout the application

}
