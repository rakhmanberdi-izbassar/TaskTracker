package main

import (
	"fmt"
	"net/http"

	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/config"
	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/handler"
	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/repository"
	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	config, err := config.NewConfig()
	if err != nil {
		return
	}

	taskRepo := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.New()

	router.GET("/api/tasks", taskHandler.GetAll)
	router.GET("/api/tasks/:id", taskHandler.GetById)
	router.DELETE("/api/tasks/:id", taskHandler.Delete)

	router.POST("/api/tasks", taskHandler.Create)
	router.PUT("/api/tasks/:id", taskHandler.Update)

	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Server error", err)
	}
}
