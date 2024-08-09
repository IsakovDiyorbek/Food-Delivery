package handler

import (
	pb "github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto"
	"github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto/user"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new notification
// @Description Create a new notification
// @Tags Notification
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param notification body pb.CreateNotificationRequest true "Notification"
// @Success 200 {object} pb.NotificationEmpty
// @Failure 400 {string} string "Bad Request"
// @Router /notification [post]
func (h *Handler) CreateNotification(c *gin.Context) {
	notification := pb.CreateNotificationRequest{}
	err := c.BindJSON(&notification)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	_, err = h.User.GetProfile(c, &user.GetProfileRequest{Id: notification.UserId})
	if err != nil{
		c.JSON(400, "User not found")
		return
	}
	_, err = h.Notification.CreateNotification(c, &notification)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Success Create Notification")
}

// @Summary Get notification by id
// @Description Get notification by id
// @Tags Notification
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Notification ID"
// @Success 200 {object} pb.Notification
// @Failure 400 {string} string "Bad Request"
// @Router /notification/{id} [get]
func (h *Handler) GetNotification(c *gin.Context) {
	id := c.Param("id")

	notification, err := h.Notification.GetNotification(c, &pb.GetNotificationRequest{Id: id})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, notification)
}

// @Router        /read [PUT]
// @Summary       MARK notification as read
// @Description   This API marks a notification as read
// @Tags          Notification
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         data body pb.MarkNotificationAsReadReq true "Notification"
// @Success       200 {object} string "message": "updated successfully"
// @Failure       400 {object} string "error": "error message"
func (h *Handler) MarkNotificationAsRead(c *gin.Context) {
	var req pb.MarkNotificationAsReadReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err := h.Notification.MarkNotificationAsRead(c, &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "updated successfully",
	})
}