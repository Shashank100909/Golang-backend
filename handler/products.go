package handlers

import (
	"net/http"
	"strconv"

	dao "github.com/Shashank100909/STUDENTS-API/daos"
	"github.com/Shashank100909/STUDENTS-API/dtos"
	"github.com/Shashank100909/STUDENTS-API/models"
	services "github.com/Shashank100909/STUDENTS-API/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productHandler struct {
	productService services.ProductService
	productDAO     dao.ProductDAO
}

func NewProductHandler(productService services.ProductService, productDAO dao.ProductDAO) *productHandler {
	return &productHandler{productService, productDAO}
}

func (h *productHandler) AddProduct(c *gin.Context) {
	var req models.Products

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.productService.AddProduct(req)
	if err != nil {
		if err.Error() == "product already exist" {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Product name already exis"})
			return
		}
		c.JSON(http.StatusInternalServerError,
			gin.H{"Error": "Internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Message": "Product added successfully"})
}

func (h *productHandler) GetProducts(c *gin.Context) {
	productIdStr := c.Query("id")
	var err error
	ProductID := 0
	if productIdStr != "" {
		ProductID, err = strconv.Atoi(productIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid product ID"})
			return
		}
	}

	Products, err := h.productService.GetProducts(ProductID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{"Error": "No products found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Products": Products,
	})
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	ProductIdstr := c.Param("id")

	ProductID, err := strconv.Atoi(ProductIdstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Inavlid product ID",
		})
		return
	}

	err = h.productService.DeleteProduct(ProductID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"Error": "Product does not exist",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Internel server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Success": "Product Deleted successfully",
	})
}

func (h *productHandler) AddProductToCart(c *gin.Context) {

	UserID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "User does not exist",
		})
	}

	// fmt.Printf("user_id ", UserID)

	var err error
	ProductID := 0
	productIdStr := c.Query("product_id")
	if productIdStr != "" {
		ProductID, err = strconv.Atoi(productIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Invalid product ID",
			})
		}
	}

	var req dtos.AddToCartReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := h.productService.AddProductToCart(UserID.(int), req, ProductID)
	if err != nil {
		if err.Error() == "Product does not exist" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Succses": "Item added to cart successfully",
		"User_id": userID,
	})
}

func (h *productHandler) GetCartItems(c *gin.Context) {

	UserID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "User does not exist",
		})
		return
	}

	cartProducts, err := h.productService.GetCartItems(UserID.(int))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"Success" :"Cart is empty",
			})
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Internal server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Success":  "Details fetched successfully",
		"Products": cartProducts,
	})
}

// func (h *productHandler) UpdateCartItems(c *gin.Context) {

// 	var req dtos.AddToCartReq
// 	UserID, exists := c.Get("user_id")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"Error": "User does not exist",
// 		})
// 	}

// 	productIdStr := c.Query("product_id")

// 	ProductID, err := strconv.Atoi(productIdStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"Error": "Invalid product ID",
// 		})
// 	}

// 	err = h.productService.UpdateCartItems(UserID.(int), ProductID, req)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"Error": err.Error(),
// 		})
// 	}
// 	c.JSON(http.StatusInternalServerError, gin.H{
// 		"Error": "Internal server error",
// 	})
// 	c.JSON(http.StatusCreated, gin.H{
// 		"Success": "Cart updated successfully",
// 	})
// }

func (h *productHandler) DeleteProductFromCart(c *gin.Context) {
    productIdStr := c.Param("product_id")
    if productIdStr == "" {
        c.JSON(http.StatusBadRequest, gin.H{"Error": "Product ID is required"})
        return
    }

    ProductID, err := strconv.Atoi(productIdStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Product ID"})
        return
    }

    UserID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"Error": "User does not exist"})
        return
    }

    err = h.productDAO.DeleteProductFromCart(ProductID, UserID.(int))
    if err == gorm.ErrRecordNotFound {
        c.JSON(http.StatusBadRequest, gin.H{"Error": "Product not found in cart"})
        return
    } else if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal server error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"Success": "Product Deleted Successfully"})
}

func (h *productHandler) AddAddress(c *gin.Context){
	 var Input models.Address

if	err := c.ShouldBindJSON(&Input); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{
		"Error" : "Invalid request payload",
	})
	return
}

UserID, exists := c.Get("user_id")
if !exists{
	c.JSON(http.StatusUnauthorized, gin.H{
		"Error":"User does not exist",
	})
	return
}

AddressID,err := h.productService.AddAddress(Input,UserID.(int))
if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{
		"Error":"Unable to add address",
	})
	return
}
c.JSON(http.StatusCreated, gin.H{
	"Succsess": "Addrress added successfully",
	"AddressID": AddressID,
})
}