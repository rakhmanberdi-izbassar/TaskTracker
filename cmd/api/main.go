package main

import (
	"fmt"
	"net/http"

	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/handler"
	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/repository"
	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	taskRepo := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.New()

	router.GET("/api/tasks", taskHandler.GetAll)
	router.GET("/api/tasks/:id", taskHandler.GetById)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Server error", err)
	}
}
