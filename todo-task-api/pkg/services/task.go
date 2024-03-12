package services

import (
	"context"
	"net/http"
	"task-svc/pkg/db"
	"task-svc/pkg/models"
	"task-svc/pkg/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedTodoServiceServer
}

func (s *Server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	var task models.Tasks

	task.Task = req.Task
	task.Description = req.Description

	if res := s.H.DB.Create(&task); res.Error != nil {
		return &pb.CreateTaskResponse{
			Status: http.StatusConflict,
			Error:  res.Error.Error(),
		}, nil
	}
	return &pb.CreateTaskResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	var task models.Tasks

	if res := s.H.DB.Delete(&task, req.Id); res.Error != nil {
		return &pb.DeleteTaskResponse{
			Status: http.StatusNotFound,
			Error:  res.Error.Error(),
		}, nil
	}
	return &pb.DeleteTaskResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) Complete(ctx context.Context, req *pb.CompleteRequest) (*pb.CompleteResponse, error) {
	var tasks models.Tasks

	if err := s.H.DB.Where("id=?", req.Id).First(&tasks).Error; err != nil {
		return &pb.CompleteResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}
	tasks.Done = true

	if err := s.H.DB.Save(&tasks).Error; err != nil {
		return &pb.CompleteResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.CompleteResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) CompleteList(ctx context.Context, req *pb.CompleteListRequest) (*pb.CompleteListResponse, error) {
	var tasks []models.Tasks

	if err := s.H.DB.Where("done=?", true).Find(&tasks).Error; err != nil {
		return &pb.CompleteListResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}
	var pbTasks []*pb.FindOne
	for _, task := range tasks {
		pbTasks = append(pbTasks, &pb.FindOne{
			Task:        task.Task,
			Description: task.Description,
			Done:        task.Done,
		})
	}

	return &pb.CompleteListResponse{
		Status: http.StatusOK,
		Data:   pbTasks,
	}, nil

}

func (s *Server) ListTask(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	var tasks []models.Tasks

	if err := s.H.DB.Find(&tasks).Error; err != nil {
		return &pb.ListResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	var pbTasks []*pb.FindOne
	for _, task := range tasks {
		pbTasks = append(pbTasks, &pb.FindOne{
			Task:        task.Task,
			Description: task.Description,
			Done:        task.Done,
		})
	}

	return &pb.ListResponse{
		Status: http.StatusOK,
		Data:   pbTasks,
	}, nil
}
