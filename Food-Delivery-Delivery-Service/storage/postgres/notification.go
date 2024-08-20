package postgres

import (
	"database/sql"
	"log"

	pb "github.com/Food-Delivery/Food-Delivery-Delivery-Service/genproto"
	"github.com/google/uuid"
)

type NotificationManager struct {
	db *sql.DB
}

func NewNotificationManager(db *sql.DB) *NotificationManager {
	return &NotificationManager{db: db}
}

func (n *NotificationManager) CreateNotification(req *pb.CreateNotificationRequest) (*pb.NotificationEmpty, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO
			notifications
				(id, 
				user_id, 
				message)
		VALUES
			($1, $2, $3)`

	_, err := n.db.Exec(
		query,
		id,
		req.UserId,
		req.Message,
	)

	if err != nil {
		return nil, err
	}
	return &pb.NotificationEmpty{}, nil
}

func (n *NotificationManager) GetNotification(req *pb.GetNotificationRequest) (*pb.Notification, error) {

	query := `
		SELECT
			id,
			user_id,
			message,
			is_read,
			created_at
		FROM
			notifications
		WHERE
			id = $1`

	row := n.db.QueryRow(query, req.Id)
	var notification pb.Notification

	err := row.Scan(
		&notification.Id,
		&notification.UserId,
		&notification.Message,
		&notification.IsRead,
		&notification.CreatedAt,
	)

	if err != nil {
		log.Println("no rows result set")
		return nil, err
	}
	return &notification, nil
}

func (n *NotificationManager) MarkNotificationAsRead(req *pb.MarkNotificationAsReadReq) (*pb.MarkNotificationAsReadResp, error) {
	query := `
		UPDATE
			notifications
		SET
			is_read = true
		WHERE
			id = $1`

	_, err := n.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.MarkNotificationAsReadResp{Success: true, Message: "Notification marked as read successfully"}, nil
}
