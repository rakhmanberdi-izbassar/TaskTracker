package handler

import (
	"net/http"
	"strconv"

	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/model"
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

func (th *TaskHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	err = th.service.Delete(id)
}

func (th *TaskHandler) Create(c *gin.Context) {
	var input model.CreateTaskInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad json data"})
		return
	}

	newId, err := th.service.Create(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Some error "})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": newId})
}

func (th *TaskHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var input model.UpdateTaskInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad json data"})
	}
	task, err := th.service.Update(id, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Some error"})
		return
	}

	c.JSON(http.StatusOK, task)
}
