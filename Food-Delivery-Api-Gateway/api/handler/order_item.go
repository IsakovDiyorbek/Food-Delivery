package handler

import (
	"fmt"
	"strconv"

	pb "github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new order item
// @Description Create a new order item
// @Tags OrderItem
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param order_item body pb.CreateOrderItemRequest true "Order Item"
// @Success 200 {object} pb.OrderItemEmpty
// @Failure 400 {string} string "Bad Request"
// @Router /order_item [post]
func (h *Handler) CreateOrderItem(c *gin.Context) {
	req := pb.CreateOrderItemRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	fmt.Println(req.OrderId, req.ProductId, req.Quantity, req.Price)
	_, err = h.OrderItems.CreateOrderItem(c, &req)


	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	_, err = h.Order.UpdateOrder(c, &pb.UpdateOrderRequest{Id: req.OrderId, Status: "Assigned"})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}


	c.JSON(200, "Success Create OrderItem")
}

// @Summary Get an order item by id
// @Description Get an order item by id
// @Tags OrderItem
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Order Item ID"
// @Success 200 {object} pb.OrderItem
// @Failure 400 {string} string "Bad Request"
// @Router /order_item/{id} [get]
func (h *Handler) GetOrderItem(c *gin.Context) {
	id := c.Param("id")
	orderItem, err := h.OrderItems.GetOrderItem(c, &pb.GetOrderItemRequest{Id: id})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, orderItem)
}

// @Summary Update an order item by id
// @Description Update an order item by id
// @Tags OrderItem
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id query string true "Order Item ID"
// @Param order_id query string false "Order ID"
// @Param product_id query string false "Product ID"
// @Param quantity query int false "Order Item quantity"
// @Param price query float64 false "Order Item price"
// @Success 200 {object} pb.OrderItemEmpty
// @Failure 400 {string} string "Bad Request"
// @Router /order_item/{id} [put]
func (h *Handler) UpdateOrderItem(c *gin.Context) {
	orderItem := pb.UpdateOrderItemRequest{}
	orderItem.Id = c.Query("id")
	orderItem.OrderId = c.Query("order_id")
	orderItem.ProductId = c.Query("product_id")
	quantityStr := c.Query("quantity")
	quantity, _ := strconv.Atoi(quantityStr)
	orderItem.Quantity = int32(quantity)

	priceStr := c.Query("price")
	if priceStr != "" {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid price parameter"})
			return
		}
		orderItem.Price = float32(price)
	}

	_, err := h.OrderItems.UpdateOrderItem(c, &orderItem)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Success Update OrderItem")
}

// @Summary Get all order items
// @Description Get all order items
// @Tags OrderItem
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param order_id query string false "Order ID"
// @Param product_id query string false "Product ID"
// @Param quantity query int false "Order Item quantity"
// @Param price query float32 false "Order Item price"
// @Success 200 {object} pb.OrderItem
// @Failure 400 {string} string "Bad Request"
// @Router /order_items [get]
func (h *Handler) GetAllOrderItems(c *gin.Context) {
	orderItem := pb.GetAllOrderItemsRequest{}
	orderItem.OrderId = c.Query("order_id")
	orderItem.ProductId = c.Query("product_id")
	
	quantityStr := c.Query("quantity")
	quantity, _ := strconv.Atoi(quantityStr)
	orderItem.Quantity = int32(quantity)
	
	priceStr := c.Query("price")
	if priceStr != "" {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid price parameter"})
			return
		}
		orderItem.Price = float32(price)
	}

	res, err := h.OrderItems.ListOrderItems(c, &orderItem)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, res)
}

// @Summary Delete an order item by id
// @Description Delete an order item by id
// @Tags OrderItem
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Order Item ID"
// @Success 200 {object} pb.OrderItemEmpty
// @Failure 400 {string} string "Bad Request"
// @Router /order_item/{id} [delete]
func (h *Handler) DeleteOrderItem(c *gin.Context) {
	id := c.Param("id")
	_, err := h.OrderItems.DeleteOrderItem(c, &pb.DeleteOrderItemRequest{Id: id})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Success Delete OrderItem")
}
