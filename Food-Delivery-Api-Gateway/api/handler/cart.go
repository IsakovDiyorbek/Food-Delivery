package handler

import (
	"strconv"

	_ "google.golang.org/protobuf/types/known/structpb"

	pb "github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto"
	"github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto/user"
	"github.com/gin-gonic/gin"
)

// CreateCart handles the creation of a new cart.
// @Summary Create a new cart
// @Description Create a new cart with the provided details.
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param cart body pb.CreateCartReq true "Cart details"
// @Success 200 {string} string "Success Create Cart"
// @Failure 400 {string} string "Error message"
// @Router /cart [post]
func (h *Handler) CreateCart(c *gin.Context) {
	req := pb.CreateCartReq{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	res, err := h.User.GetProfile(c, &user.GetProfileRequest{Id: req.UserId})
	if err != nil {
		c.JSON(400, "User not found")
		return
	}
	if res.Role != "user" {
		c.JSON(400, "User not found")
		return
	}
	_, err = h.Cart.CreateCart(c, &req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, "Success Create Cart")
}

// @Summary Get a cart by id
// @Description Get a cart by id
// @Tags Cart
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id query string true "Cart ID"
// @Success 200 {object} pb.Cart
// @Failure 400 {string} string "Bad Request"
// @Router /cart/{id} [get]
func (h *Handler) GetCart(c *gin.Context) {
	id := c.Query("id")
	cart, err := h.Cart.GetCart(c, &pb.GetByIdCartRequest{Id: id})
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, cart)

}

// @Summary Get all carts
// @Description Get all carts
// @Tags Cart
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param quantity query int false "Cart quantity"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.Cart
// @Failure 400 {string} string "Bad Request"
// @Router /carts [get]
func (h *Handler) GetAllCarts(c *gin.Context) {
	req := pb.GetAllCartsReq{}
	quantityStr := c.Query("quantity")
	quantity, _ := strconv.Atoi(quantityStr)
	req.Quantity = int32(quantity)
	limit := c.Query("limit")
	offset := c.Query("offset")
	if limit != "" {
		intlimit, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		req.Limit = int32(intlimit)
	}
	if offset != "" {
		intoffset, err := strconv.Atoi(offset)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		req.Offset = int32(intoffset)
	}
	res, err := h.Cart.GetAllCarts(c, &req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, res)
}

// @Summary Update a cart by id
// @Description Update a cart by id
// @Tags Cart
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id query string true "Cart ID"
// @Param user_id query string false "Cart user id"
// @Param product_id query string false "Cart product id"
// @Param options query string false "Cart options"
// @Param quantity query int false "Cart quantity"
// @Success 200 {object} pb.UpdateCartRes
// @Failure 400 {string} string "Bad Request"
// @Router /cart/{id} [put]
func (h *Handler) UpdateCart(c *gin.Context) {
	req := pb.UpdateCartReq{}
	req.Id = c.Query("id")
	req.UserId = c.Query("user_id")
	req.Product = c.Query("product_id")
	req.Options = c.Query("options")
	quantityStr := c.Query("quantity")
	quantity, _ := strconv.Atoi(quantityStr)
	req.Quantity = int32(quantity)

	_, err := h.Cart.UpdateCart(c, &req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, "Success Update Cart")
}

// @Summary Delete a cart by id
// @Description Delete a cart by id
// @Tags Cart
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id query string true "Cart ID"
// @Success 200 {object} pb.DeleteCartResp
// @Failure 400 {string} string "Bad Request"
// @Router /cart/{id} [delete]
func (h *Handler) DeleteCart(c *gin.Context) {
	req := pb.DeleteCartRequest{}
	req.Id = c.Query("id")
	_, err := h.Cart.DeleteCart(c, &req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, "Success Delete Cart")
}
