package handler

import (
	"strconv"

	pb "github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto"
	"github.com/gin-gonic/gin"
)

// @Summary Create a new product
// @Description Create a new product
// @Tags Product
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param borrower body pb.CreateProductRequest true "Borrower"
// @Success 200 {object} pb.ProductEmpty
// @Failure 400 {string} string "Bad Request"
// @Router /product [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	product := pb.CreateProductRequest{}
	err := c.BindJSON(&product)

	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	_, err = h.Product.CreateProduct(c, &product)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Success Create Product")
}

// @Summary Get a product by id
// @Description Get a product by id
// @Tags Product
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Success 200 {object} pb.Product
// @Failure 400 {string} string "Bad Request"
// @Router /product/{id} [get]
func (h *Handler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := h.Product.GetProduct(c, &pb.GetProductRequest{Id: id})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, product)
}

// @Summary Update a product by id
// @Description Update a product by id
// @Tags Product
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id query string true "Product ID"
// @Param name query string false "Product name"
// @Param image_url query string false "Product image url"
// @Param description query string false "Product description"
// @Param price query float64 false "Product price"
// @Success 200 {object} pb.ProductEmpty
// @Failure 400 {string} string "Bad Request"
// @Router /product/{id} [put]
func (h *Handler) UpdateProduct(c *gin.Context) {
	product := pb.UpdateProductRequest{}
	product.Id = c.Query("id")
	product.Description = c.Query("description")
	product.ImageUrl = c.Query("image_url")
	product.Name = c.Query("name")
	priceStr := c.Query("price")
	if priceStr != "" {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid price parameter"})
			return
		}
		product.Price = float32(price)
	}


	_, err := h.Product.UpdateProduct(c, &product)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Success Update Product")
}


// @Summary Delete a product by id
// @Description Delete a product by id
// @Tags Product
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Success 200 {object} pb.ProductEmpty
// @Failure 400 {string} string "Bad Request"
// @Router /product/{id} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	_, err := h.Product.DeleteProduct(c, &pb.DeleteProductRequest{Id: id})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "Success Delete Product")
}

// @Summary Get all products
// @Description Get all products
// @Tags Product
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param name query string false "Product name"
// @Param image_url query string false "Product image url"
// @Param description query string false "Product description"
// @Param price query float64 false "Product price"
// @Success 200 {object} pb.Product
// @Failure 400 {string} string "Bad Request"
// @Router /products [get]
func (h *Handler) GetAllProducts(c *gin.Context) {
	req := pb.GetAllProductsRequest{}
	req.Name = c.Query("name")
	req.ImageUrl = c.Query("image_url")
	req.Description = c.Query("description")
	priceStr := c.Query("price")
	if priceStr != "" {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid price parameter"})
			return
		}
		req.Price = float32(price)
	}

	res, err := h.Product.ListProducts(c, &req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, res)
}
