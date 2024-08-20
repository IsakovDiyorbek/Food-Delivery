package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/Food-Delivery/Food-Delivery-Delivery-Service/genproto"
	"github.com/google/uuid"
)

type CourierLocationRepo struct {
	db *sql.DB
}

func NewCourierLocationRepo(db *sql.DB) *CourierLocationRepo {
	return &CourierLocationRepo{db: db}
}

func (c *CourierLocationRepo) CreateCourierLocation(req *pb.CreateCourierLocationRequest) (*pb.Empty, error) {
	query := `insert into courier_locations(id, courier_id, latitude, longitude, start_time, end_time, status) values($1, $2, $3, $4, $5, $6, $7)`

	_, err := c.db.Exec(query, uuid.NewString(), req.CourierId, req.Latitude, req.Longitude, req.StartTime, req.EndTime, req.Status)
	if err != nil {
		log.Fatal("Error while creating courier location", err)
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (c *CourierLocationRepo) GetCourierLocation(req *pb.GetCourierLocationRequest) (*pb.CourierLocation, error) {
	query := `SELECT id, courier_id, latitude, longitude, start_time, end_time, status FROM courier_locations WHERE id = $1 and deleted_at = 0`

	row := c.db.QueryRow(query, req.Id)

	var courierLocation pb.CourierLocation
	err := row.Scan(
		&courierLocation.Id,
		&courierLocation.CourierId,
		&courierLocation.Latitude,
		&courierLocation.Longitude,
		&courierLocation.StartTime,
		&courierLocation.EndTime,
		&courierLocation.Status,
	)
	if err != nil {
		log.Fatal("Error while scan courier location", err)
		return nil, err
	}
	return &courierLocation, nil
}

func (c *CourierLocationRepo) UpdateCourierLocation(req *pb.UpdateCourierLocationRequest) (*pb.Empty, error) {
	var args []interface{}
	var conditions []string

	if req.CourierId != "" && req.CourierId != "string" {
		args = append(args, req.CourierId)
		conditions = append(conditions, fmt.Sprintf("courier_id = $%d", len(args)))
	}

	if req.Latitude != 0 {
		args = append(args, req.Latitude)
		conditions = append(conditions, fmt.Sprintf("latitude = $%d", len(args)))
	}

	if req.Longitude != 0 {
		args = append(args, req.Longitude)
		conditions = append(conditions, fmt.Sprintf("longitude = $%d", len(args)))
	}

	if req.StartTime != "" && req.StartTime != "string" {
		args = append(args, req.StartTime)
		conditions = append(conditions, fmt.Sprintf("start_time = $%d", len(args)))
	}

	if req.EndTime != "" && req.EndTime != "string" {
		args = append(args, req.EndTime)
		conditions = append(conditions, fmt.Sprintf("end_time = $%d", len(args)))
	}
	if req.Status != "" {
		args = append(args, req.Status)
		conditions = append(conditions, fmt.Sprintf("status = $%d", len(args)))
	}
	args = append(args, time.Now())
	conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)))

	query := `UPDATE courier_locations SET ` + strings.Join(conditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(args)+1)
	args = append(args, req.Id)

	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (c *CourierLocationRepo) ListCourierLocations(req *pb.GetAllCourierLocationsRequest) (*pb.CourierLocationList, error) {
	query := `
		SELECT
			id, 
			courier_id, 
			latitude, 
			longitude, 
			start_time, 
			end_time, 
			status
		FROM
			courier_locations
		`

	var args []interface{}
	argCount := 1
	filters := []string{}

	if req.CourierId != "" {
		filters = append(filters, fmt.Sprintf("courier_id = $%d", argCount))
		args = append(args, req.CourierId)
		argCount++
	}

	if req.Latitude!= 0 {
		filters = append(filters, fmt.Sprintf("latitude = $%d", argCount))
		args = append(args, req.Latitude)
		argCount++
	}

	if req.Longitude != 0 {
		filters = append(filters, fmt.Sprintf("longitude = $%d", argCount))
		args = append(args, req.Longitude)
		argCount++
	}

	if req.StartTime != "" {
		filters = append(filters, fmt.Sprintf("start_time = $%d", argCount))
		args = append(args, req.StartTime)
		argCount++
	}

	if req.EndTime != "" {
		filters = append(filters, fmt.Sprintf("end_time = $%d", argCount))
		args = append(args, req.EndTime)
		argCount++
	}

	if req.Status != "" {
		filters = append(filters, fmt.Sprintf("status = $%d", argCount))
		args = append(args, req.Status)
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

	rows, err := c.db.Query(query, args...)
	if err != nil {
		log.Println("no rows result set")
		return nil, err
	}
	defer rows.Close()

	var courierLocations []*pb.CourierLocation
	for rows.Next() {
		var courierLocation pb.CourierLocation
		err := rows.Scan(
			&courierLocation.Id,
			&courierLocation.CourierId,
			&courierLocation.Latitude,
			&courierLocation.Longitude,
			&courierLocation.StartTime,
			&courierLocation.EndTime,
			&courierLocation.Status,
		)
		if err != nil {
			log.Println("no rows result set")
			return nil, err
		}
		courierLocations = append(courierLocations, &courierLocation)
	}

	return &pb.CourierLocationList{CourierLocations: courierLocations}, nil

}

func (c *CourierLocationRepo) DeleteCourierLocation(req *pb.DeleteCourierLocationRequest) (*pb.Empty, error) {
	query := `
	UPDATE
		courier_locations
	SET
		deleted_at = extract(epoch from now())
	WHERE
		id = $1
	`

	_, err := c.db.Exec(query, req.Id)

	if err != nil {
		log.Fatal("Error while deleting courier location", err)
		return nil, err
	}
	return &pb.Empty{}, nil

}

