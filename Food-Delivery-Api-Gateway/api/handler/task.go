package handler

import (
	pb "github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new task
// @Description Create a new task
// @Tags Task
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param task body pb.CreateTaskRequest true "Task"
// @Success 200 {object} pb.TaskEmpty
// @Failure 400 {string} string "Bad Request"
// @Router /task [post]
func (h *Handler) CreateTask(c *gin.Context) {
	task := pb.CreateTaskRequest{}
	err := c.BindJSON(&task)

	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	_, err = h.Task.CreateTask(c, &task)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Success Create Task")
}

// @Summary Get a task by id
// @Description Get a task by id
// @Tags Task
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id query string true "Task ID"
// @Success 200 {object} pb.Task
// @Failure 400 {string} string "Bad Request"
// @Router /task/{id} [get]
func (h *Handler) GetTask(c *gin.Context) {
	id := c.Query("id")
	task, err := h.Task.GetTask(c, &pb.GetTaskRequest{Id: id})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, task)
}

// @Summary Update a task by id
// @Description Update a task by id
// @Tags Task
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id query string true "Task ID"
// @Param title query string false "Task title"
// @Param description query string false "Task description"
// @Param user_id_assigned_to query string false "Task assigned to"
// @Param status query string false "Task status"
// @Param due_date query string false "Task due date"
// @Success 200 {object} pb.TaskEmpty
// @Failure 400 {string} string "Bad Request"
// @Router /task/{id} [put]
func (h *Handler) UpdateTask(c *gin.Context) {
	task := pb.UpdateTaskRequest{}
	task.Id = c.Query("id")
	task.Description = c.Query("description")
	task.Title = c.Query("title")
	task.AssignedTo = c.Query("user_id_assigned_to")
	task.Status = c.Query("status")
	task.DueDate = c.Query("due_date")

	_, err := h.Task.UpdateTask(c, &task)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Success Update Task")
}

// @Summary Get all tasks
// @Description Get all tasks
// @Tags Task
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param title query string false "Task title"
// @Param description query string false "Task description"
// @Param assigned_to query string false "Task assigned to"
// @Param status query string false "Task status"
// @Param due_date query string false "Task due date"
// @Success 200 {object} pb.Task
// @Failure 400 {string} string "Bad Request"
// @Router /tasks [get]
func (h *Handler) GetAllTasks(c *gin.Context) {
	task := pb.GetAllTasksRequest{}
	task.Title = c.Query("title")
	task.Description = c.Query("description")
	task.AssignedTo = c.Query("assigned_to")
	task.Status = c.Query("status")
	task.DueDate = c.Query("due_date")

	res, err := h.Task.ListTasks(c, &task)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, res)
}


// @Summary Delete a task by id
// @Description Delete a task by id
// @Tags Task
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id query string true "Task ID"
// @Success 200 {object} pb.TaskEmpty
// @Failure 400 {string} string "Bad Request"
// @Router /task{id} [delete]
func (h *Handler) DeleteTask(c *gin.Context) {
	id := c.Query("id")
	_, err := h.Task.DeleteTask(c, &pb.DeleteTaskRequest{Id: id})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Success Delete Task")
}
