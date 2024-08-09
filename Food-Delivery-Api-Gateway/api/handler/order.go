package handler

import (
	"net/http"
	"strconv"

	pb "github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto"
	"github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto/user"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new order
// @Description Create a new order
// @Tags Order
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param order body pb.CreateOrderRequest true "Order"
// @Success 200 {object} pb.OrderEmpty
// @Failure 400 {string} string "Bad Request"
// @Router /order [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	req := pb.CreateOrderRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	resUser, err := h.User.GetProfile(c, &user.GetProfileRequest{Id: req.UserId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	resCourier, err := h.User.GetProfile(c, &user.GetProfileRequest{Id: req.CourierId})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Courier not found"})
		return
	}

	if resUser.Role != "user" || resCourier.Role != "courier" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user or courier role"})
		return
	}

	_, err = h.Order.CreateOrder(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, "Order created successfully")
}

// @Summary Get an order by id
// @Description Get an order by id
// @Tags Order
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Success 200 {object} pb.Order
// @Failure 400 {string} string "Bad Request"
// @Router /order/{id} [get]
func (h *Handler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := h.Order.GetOrder(c, &pb.GetOrderRequest{Id: id})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, order)
}

// @Summary Update an order by id
// @Description Update an order by id
// @Tags Order
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id query string true "Order ID"
// @Param status query string false "Order status"
// @Param courier_id query string false "Order courier id"
// @Param delivery_address query string false "Order delivery address"
// @Param delivery_time query string false "Order delivery time"
// @Success 200 {object} pb.OrderEmpty
// @Failure 400 {string} string "Bad Request"
// @Router /order/{id} [put]
func (h *Handler) UpdateOrder(c *gin.Context) {
	order := pb.UpdateOrderRequest{}
	order.Id = c.Query("id")
	order.UserId = c.Query("user_id")
	order.Status = c.Query("status")
	order.CourierId = c.Query("courier_id")
	order.DeliveryAddress = c.Query("delivery_address")
	totalAmountStr := c.Query("total_amount")
	if totalAmountStr != "" {
		totalAmount, err := strconv.ParseFloat(totalAmountStr, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid total_amount parameter"})
			return
		}
		order.TotalAmount = float32(totalAmount)
	}

	_, err := h.Order.UpdateOrder(c, &order)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Success Update Order")
}

// @Summary Get all orders
// @Description Get all orders
// @Tags Order
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param user_id query string false "Order user id"
// @Param status query string false "Order status"
// @Param courier_id query string false "Order courier id"
// @Param delivery_address query string false "Order delivery address"
// @Param delivery_time query string false "Order delivery time"
// @Success 200 {object} pb.Order
// @Failure 400 {string} string "Bad Request"
// @Router /orders [get]
func (h *Handler) GetAllOrders(c *gin.Context) {
	req := pb.GetAllOrdersRequest{}
	req.UserId = c.Query("user_id")
	req.Status = c.Query("status")
	req.CourierId = c.Query("courier_id")
	req.DeliveryAddress = c.Query("delivery_address")

	totalAmountStr := c.Query("total_amount")
	if totalAmountStr != "" {
		totalAmount, err := strconv.ParseFloat(totalAmountStr, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid total_amount parameter"})
			return
		}
		req.TotalAmount = float32(totalAmount)
	}

	res, err := h.Order.ListOrders(c, &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}

// @Summary Delete an order by id
// @Description Delete an order by id
// @Tags Order
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Success 200 {object} pb.OrderEmpty
// @Failure 400 {string} string "Bad Request"
// @Router /order/{id} [delete]
func (h *Handler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	_, err := h.Order.DeleteOrder(c, &pb.DeleteOrderRequest{Id: id})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Success Delete Order")
}
