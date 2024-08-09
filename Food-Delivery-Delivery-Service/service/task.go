package service

import (
	"context"

	"github.com/Food/Food-Delivery-Delivery-Service/genproto"
	"github.com/Food/Food-Delivery-Delivery-Service/storage"
)

type TaskService struct {
	storage storage.StorageI
	genproto.UnimplementedTaskServiceServer
}

func NewTaskService(storage storage.StorageI) *TaskService {
	return &TaskService{storage: storage}
}

func (t *TaskService) CreateTask(ctx context.Context, req *genproto.CreateTaskRequest) (*genproto.TaskEmpty, error) {
	return t.storage.Task().CreateTask(req)
}

func (t *TaskService) GetTask(ctx context.Context, req *genproto.GetTaskRequest) (*genproto.Task, error) {
	return t.storage.Task().GetTask(req)
}

func (t *TaskService) UpdateTask(ctx context.Context, req *genproto.UpdateTaskRequest) (*genproto.TaskEmpty, error) {
	return t.storage.Task().UpdateTask(req)
}

func (t *TaskService) DeleteTask(ctx context.Context, req *genproto.DeleteTaskRequest) (*genproto.TaskEmpty, error) {
	return t.storage.Task().DeleteTask(req)
}

func (t *TaskService) ListTasks(ctx context.Context, req *genproto.GetAllTasksRequest) (*genproto.TaskList, error) {
	return t.storage.Task().ListTasks(req)
}
