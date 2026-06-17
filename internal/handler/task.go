package handler

import (
	"net/http"
	"strconv"

	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(taskSrv *service.TaskService) *TaskHandler {
	return &TaskHandler{service: taskSrv}
}

func (th *TaskHandler) GetAll(c *gin.Context) {
	tasks, err := th.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (th *TaskHandler) GetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	task, err := th.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, task)
}
