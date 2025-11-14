package app

import (
	"hrm-app/config"
	"hrm-app/internal/domain/admin"
	"hrm-app/internal/domain/auth"
	"hrm-app/internal/domain/department"
	"hrm-app/internal/domain/employee"
	"hrm-app/internal/domain/manager"
	"hrm-app/internal/domain/presence"
	"hrm-app/internal/domain/user"
	"hrm-app/internal/domain/work_hour"
	"hrm-app/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})

		// User routes
		userRepo := user.NewRepository()
		userUseCase := user.NewUseCase(userRepo)
		userHandler := user.NewHandler(userUseCase)

		// Employee routes
		employeeRepo := employee.NewRepository()
		employeeUsecase := employee.NewUseCase(employeeRepo)
		employeeHandler := employee.NewHandler(employeeUsecase)

		// Manager routes
		managerRepo := manager.NewRepository()
		managerUseCase := manager.NewUseCase(managerRepo)
		managerHandler := manager.NewHandler(managerUseCase)

		// Department routes
		departmentRepo := department.NewRepository()
		departmentUseCase := department.NewUseCase(departmentRepo)
		departmentHandler := department.NewHandler(departmentUseCase)

		// Admin routes
		adminRepo := admin.NewRepository()
		adminUseCase := admin.NewUseCase(adminRepo)
		adminHandler := admin.NewHandler(adminUseCase)

		// Work Hour routes
		workHourRepo := work_hour.NewRepository()

		// Presence routes
		presenceRepo := presence.NewRepository()
		presenceUseCase := presence.NewUseCase(presenceRepo, employeeRepo, workHourRepo)
		presenceHandler := presence.NewHandler(presenceUseCase)


		// auth handler needs repo + cfg
		authHandler := auth.NewHandler(userRepo, cfg)

		auth := r.Group("/api/presence")
		auth.Use(middleware.AuthMiddleware(cfg))
		{
			auth.POST("/checkin", presenceHandler.Checkin)
			auth.PUT("/checkout", presenceHandler.Checkout)
		}

		api.POST("/login", authHandler.Login)

		user := api.Group("/users")
		{
			user.POST("/", userHandler.Register)
			user.GET("/", userHandler.GetAll)
			user.GET("/:id", userHandler.GetByID)
			user.DELETE("/:id", userHandler.Delete)
		}
		employee := api.Group("/employees")
		{
			employee.POST("/", employeeHandler.RegisterWithContact)
			employee.GET("/", employeeHandler.GetAll)
			employee.GET("/:id", employeeHandler.GetByID)
			employee.DELETE("/:id", employeeHandler.Delete)
		}
		manager := api.Group("/managers")
		{
			manager.POST("/", managerHandler.RegisterWithContact)
			manager.GET("/", managerHandler.GetAll)
			manager.GET("/:id", managerHandler.GetByID)
			manager.DELETE("/:id", managerHandler.Delete)
		}
		department := api.Group("/departments")
		{
			department.POST("/", departmentHandler.Register)
			department.GET("/", departmentHandler.GetAll)
			department.GET("slug/:slug", departmentHandler.GetBySlug)
			department.GET("/:id", departmentHandler.GetByID)
			department.DELETE("/:id", departmentHandler.Delete)
		}
		admin := api.Group("/admins")
		{
			admin.POST("/", adminHandler.RegisterWithContact)
			admin.GET("/", adminHandler.GetAll)
			admin.GET("/:id", adminHandler.GetByID)
			admin.DELETE("/:id", adminHandler.Delete)
		}
	}

	return r
}
