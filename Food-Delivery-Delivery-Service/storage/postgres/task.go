package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/Food/Food-Delivery-Delivery-Service/genproto"
	"github.com/google/uuid"
)

type TasksRepo struct {
	db *sql.DB
}

func NewTasksRepo(db *sql.DB) *TasksRepo {
	return &TasksRepo{db: db}
}

func (t *TasksRepo) CreateTask(req *pb.CreateTaskRequest) (*pb.TaskEmpty, error) {
	query := `insert into tasks(id, title, description, status, user_id_assigned_to, due_date) values($1, $2, $3, $4, $5, $6)`

	id := uuid.NewString()
	_, err := t.db.Exec(query, id, req.Title, req.Description, req.Status, req.AssignedTo, req.DueDate)
	if err != nil {
		log.Fatal("Error while creating task", err)
		return nil, err
	}
	return &pb.TaskEmpty{}, nil
}

func (t *TasksRepo) GetTask(req *pb.GetTaskRequest) (*pb.Task, error) {
	query := `
	SELECT
		id, 
		title, 
		description, 
		status, 
		user_id_assigned_to, 
		due_date,
		created_at
	FROM
		tasks
	WHERE
		id = $1
	AND 
		deleted_at = 0
	`

	row := t.db.QueryRow(query, req.Id)
	var task pb.Task
	err := row.Scan(
		&task.Id, &task.Title, &task.Description, &task.Status, &task.AssignedTo, &task.DueDate, &task.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No task found with id %v", req.Id)
			return nil, fmt.Errorf("no task found with id %v", req.Id)
		}
		log.Printf("Error while scanning task: %v", err)
		return nil, err
	}
	return &task, nil
}


func (t *TasksRepo) UpdateTask(req *pb.UpdateTaskRequest) (*pb.TaskEmpty, error) {
	var args []interface{}
	var conditions []string

	if req.Title != "" && req.Title != "string" {
		args = append(args, req.Title)
		conditions = append(conditions, fmt.Sprintf("title = $%d", len(args)))
	}
	if req.Description != "" && req.Description != "string" {
		args = append(args, req.Description)
		conditions = append(conditions, fmt.Sprintf("description = $%d", len(args)))
	}
	if req.Status != "" {
		args = append(args, req.Status)
		conditions = append(conditions, fmt.Sprintf("status = $%d", len(args)))
	}
	if req.AssignedTo != "" {
		args = append(args, req.AssignedTo)
		conditions = append(conditions, fmt.Sprintf("user_id_assigned_to = $%d", len(args)))
	}

	if req.DueDate != "" {
		args = append(args, req.DueDate)
		conditions = append(conditions, fmt.Sprintf("due_date = $%d", len(args)))
	}
	args = append(args, time.Now())
	conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)))

	query := `UPDATE tasks SET ` + strings.Join(conditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(args)+1)
	args = append(args, req.Id)

	_, err := t.db.Exec(query, args...)
	if err != nil {
		log.Fatal("Error while updating task", err)
		return nil, err
	}

	return &pb.TaskEmpty{}, nil

}

func (t *TasksRepo) ListTasks(req *pb.GetAllTasksRequest) (*pb.TaskList, error) {
	query := `
		SELECT
			id, 
			title, 
			description, 
			status, 
			user_id_assigned_to, 
			due_date,
			created_at
		FROM
			tasks
		`

	var args []interface{}
	argCount := 1
	filters := []string{}

	if req.Title != "" {
		filters = append(filters, fmt.Sprintf("title = $%d", argCount))
		args = append(args, req.Title)
		argCount++
	}

	if req.Description != "" {
		filters = append(filters, fmt.Sprintf("description = $%d", argCount))
		args = append(args, req.Description)
		argCount++
	}

	if req.Status != "" {
		filters = append(filters, fmt.Sprintf("status = $%d", argCount))
		args = append(args, req.Status)
		argCount++
	}

	if req.AssignedTo != "" {
		filters = append(filters, fmt.Sprintf("user_id_assigned_to = $%d", argCount))
		args = append(args, req.AssignedTo)
		argCount++
	}

	if req.DueDate != "" {
		filters = append(filters, fmt.Sprintf("due_date = $%d", argCount))
		args = append(args, req.DueDate)
		argCount++
	}

	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}

	if req.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argCount)
		args = append(args, req.Limit)
		argCount++

		if req.Offset > 0 {
			query += fmt.Sprintf(" OFFSET $%d", argCount)
			args = append(args, req.Offset)
			argCount++
		}
	}

	rows, err := t.db.Query(query, args...)
	if err != nil {
		log.Println("no rows result set")
		return nil, err
	}
	defer rows.Close()

	var tasks []*pb.Task
	for rows.Next() {
		var task pb.Task
		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.UserId,
			&task.DueDate,
			&task.CreatedAt,
		)
		if err != nil {
			log.Println("no rows result set")
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return &pb.TaskList{Tasks: tasks}, nil

}

func (t *TasksRepo) DeleteTask(req *pb.DeleteTaskRequest) (*pb.TaskEmpty, error) {
	query := `
	UPDATE
		tasks
	SET
		deleted_at = extract(epoch from now())
	WHERE
		id = $1
	`

	_, err := t.db.Exec(query, req.Id)

	if err != nil {
		log.Fatal("Error while deleting task", err)
		return nil, err
	}
	return &pb.TaskEmpty{}, nil

}
