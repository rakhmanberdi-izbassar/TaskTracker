package handler

import (
	"net/http"
	"strconv"

	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/apperrors"
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
		respondWithError(c, apperrors.BadRequest("invalid ID parameter", nil))
		return
	}
	task, err := th.service.GetById(id)
	if err != nil {
		respondWithError(c, apperrors.NotFound("task not found", nil))
	}
	c.JSON(http.StatusOK, task)
}

func (th *TaskHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid ID parameter", nil))
		return
	}
	err = th.service.Delete(id)
}

func (th *TaskHandler) Create(c *gin.Context) {
	var input model.CreateTaskInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid input body", err))
		return
	}

	newId, err := th.service.Create(input)

	if err != nil {
		respondWithError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": newId})
}

func (th *TaskHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid ID parameter", nil))
		return
	}

	var input model.UpdateTaskInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid json data", nil))
	}
	updatedId, err := th.service.Update(id, input)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": updatedId})
}
