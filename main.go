package main

import (
	"fmt"

	"github.com/leo-andrei/logging/config"
	"github.com/leo-andrei/logging/drivers"
	"github.com/leo-andrei/logging/log"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	var driver log.Driver
	switch cfg.DriverType {
	case "cli":
		driver = drivers.CLIDriver{}
	case "file":
		driver = drivers.FileDriver{FilePath: cfg.FilePath}
	default:
		fmt.Println("Unknown driver type")
		return
	}

	logger := log.NewLogger(driver)

	logger.Log(log.DEBUG, "This is a debug message", map[string]interface{}{
		"CustomerId":  "12345",
		"Environment": "production",
	})

	txn := log.NewTransaction("trace-1", logger, map[string]interface{}{
		"User": "john_doe",
	})
	txn.Log(log.ERROR, "Transaction error occurred")
}
