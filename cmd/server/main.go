package main

import (
	"fmt"
	"hrm-app/config"
	"hrm-app/internal/app"
	"hrm-app/internal/domain/employee"
	"hrm-app/internal/pkg/database"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectDatabase(cfg)
	database.ConnectRedis(cfg)

	r := app.SetupRouter(cfg)
	port := fmt.Sprintf(":%d", cfg.Server.Port)
	// Auto migrate database schemas
	database.DB.AutoMigrate(&employee.Employee{})
	r.Run(port)
}