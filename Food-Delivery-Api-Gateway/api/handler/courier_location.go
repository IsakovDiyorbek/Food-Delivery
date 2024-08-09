package handler

import (
	"strconv"

	pb "github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto"
	"github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto/user"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new courier location
// @Description Create a new courier location
// @Tags CourierLocation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param courier_location body pb.CreateCourierLocationRequest true "Courier Location"
// @Success 200 {object} pb.Empty
// @Failure 400 {string} string "Bad Request"
// @Router /courier_location [post]
func (h *Handler) CreateCourierLocation(c *gin.Context) {
	reqp := pb.CreateCourierLocationRequest{}
	err := c.BindJSON(&reqp)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	res, err := h.User.GetProfile(c, &user.GetProfileRequest{Id: reqp.CourierId})
	if res.Role == "courier"{
		c.JSON(400, "User not found")
		return
	}
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	_, err = h.CourierLocation.CreateCourierLocation(c, &reqp)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Success Create Courier Location")
}

// @Summary Get a courier location by id
// @Description Get a courier location by id
// @Tags CourierLocation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id query string true "Courier Location ID"
// @Success 200 {object} pb.CourierLocation
// @Failure 400 {string} string "Bad Request"
// @Router /courier_location/{id} [get]	
func (h *Handler) GetCourierLocation(c *gin.Context) {
	reqp := pb.GetCourierLocationRequest{}
	reqp.Id = c.Query("id")
	res, err := h.CourierLocation.GetCourierLocation(c, &reqp)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, res)
}

// @Summary Update a courier location by id
// @Description Update a courier location by id
// @Tags CourierLocation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id query string true "Courier Location ID"
// @Param courier_id query string false "Courier ID"
// @Param latitude query float64 false "Courier Location latitude"
// @Param longitude query float64 false "Courier Location longitude"
// @Param start_time query string false "Courier Location start time"
// @Param end_time query string false "Courier Location end time"
// @Param status query string false "Courier Location status"
// @Success 200 {object} pb.Empty
// @Failure 400 {string} string "Bad Request"
// @Router /courier_location/{id} [put]
func (h *Handler) UpdateCourierLocation(c *gin.Context) {
	reqp := pb.UpdateCourierLocationRequest{}
	reqp.Id = c.Query("id")
	reqp.CourierId = c.Query("courier_id")
	
	latitudeStr := c.Query("latitude")
	if latitudeStr != "" {
		latitude, err := strconv.ParseFloat(latitudeStr, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid latitude parameter"})
			return
		}
		reqp.Latitude = float32(latitude)
	}

	longitudeStr := c.Query("longitude")
	if longitudeStr != "" {
		longitude, err := strconv.ParseFloat(longitudeStr, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid price parameter"})
			return
		}
		reqp.Latitude = float32(longitude)
	}

	reqp.StartTime = c.Query("start_time")
	reqp.EndTime = c.Query("end_time")
	reqp.Status = c.Query("status")

	_, err := h.CourierLocation.UpdateCourierLocation(c, &reqp)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Success Update Courier Location")
}


// @Summary Delete a courier location by id
// @Description Delete a courier location by id
// @Tags CourierLocation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id query string true "Courier Location ID"
// @Success 200 {object} pb.Empty
// @Failure 400 {string} string "Bad Request"
// @Router /courier_location{id} [delete]
func (h *Handler) DeleteCourierLocation(c *gin.Context) {
	reqp := pb.DeleteCourierLocationRequest{}
	reqp.Id = c.Query("id")
	_, err := h.CourierLocation.DeleteCourierLocation(c, &reqp)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Success Delete Courier Location")
}


// @Summary Get all courier locations
// @Description Get all courier locations
// @Tags CourierLocation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param courier_id query string false "Courier ID"
// @Param start_time query string false "Courier Location start time"
// @Param end_time query string false "Courier Location end time"
// @Param status query string false "Courier Location status"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.CourierLocation
// @Failure 400 {string} string "Bad Request"
// @Router /courier_locations [get]
func (h *Handler) GetAllCourierLocations(c *gin.Context) {
	reqp := pb.GetAllCourierLocationsRequest{}
	reqp.CourierId = c.Query("courier_id")
	reqp.StartTime = c.Query("start_time")
	reqp.EndTime = c.Query("end_time")
	reqp.Status = c.Query("status")

	if reqp.Limit != 0 && reqp.Offset != 0{
		intlimit, err := strconv.Atoi(c.Query("limit"))
		if err != nil{
			c.JSON(400, err.Error())
			return
		}
		reqp.Limit = int32(intlimit)
		offset, err := strconv.Atoi(c.Query("offset"))
		if err != nil{
			c.JSON(400, err.Error())
			return
		}
		reqp.Offset = int32(offset)
	
	}

	res, err := h.CourierLocation.ListCourierLocations(c, &reqp)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, res)
}
