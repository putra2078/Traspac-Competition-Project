package app

import (
	"hrm-app/config"
	"hrm-app/internal/domain/auth"
	"hrm-app/internal/domain/employee"
	"hrm-app/internal/domain/user"

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

		// auth handler needs repo + cfg
		authHandler := auth.NewHandler(userRepo, cfg)

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
	}

	return r
}
