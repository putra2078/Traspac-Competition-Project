package main

import (
	"fmt"
	"hrm-app/config"
	"hrm-app/internal/app"
	"hrm-app/internal/domain/contact"
	"hrm-app/internal/domain/employee"
	"hrm-app/internal/domain/manager"
	"hrm-app/internal/domain/user"
	"hrm-app/internal/pkg/database"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectDatabase(cfg)
	database.ConnectRedis(cfg)

	r := app.SetupRouter(cfg)
	port := fmt.Sprintf(":%d", cfg.Server.Port)
	// Auto migrate database schemas
	// ensure contact table exists as we now create contacts in a transaction
	database.DB.AutoMigrate(&employee.Employee{}, &contact.Contact{}, &user.User{}, &manager.Manager{})
	r.Run(port)
}
