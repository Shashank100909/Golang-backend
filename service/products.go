package services

import (
	dao "github.com/Shashank100909/STUDENTS-API/daos"
	"github.com/Shashank100909/STUDENTS-API/dtos"
	"github.com/Shashank100909/STUDENTS-API/models"
)

type ProductService interface {
	AddProduct(req models.Products) error
	GetProducts(ProductID int) ([]models.Products, error)
	DeleteProduct(ProductID int) error

	AddProductToCart(UserID int, req dtos.AddToCartReq, ProductID int) (int, error)
	GetCartItems(UserID int) ([]dtos.CartResp, error)
	AddAddress(Input models.Address, UserID int) (int, error)
	// UpdateCartItems(UserID int, ProductID int, req dtos.AddToCartReq) error
}

type productService struct {
	productDAO dao.ProductDAO
}

func NewProductService(productDAO dao.ProductDAO) ProductService {
	return &productService{productDAO}
}

func (h *productService) AddProduct(req models.Products) error {
	product := models.Products{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	}

	err := h.productDAO.AddProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (h *productService) GetProducts(ProductID int) ([]models.Products, error) {
	return h.productDAO.GetProducts(ProductID)
}

func (h *productService) DeleteProduct(ProductID int) error {
	return h.productDAO.DeleteProduct(ProductID)
}
func (h *productService) AddProductToCart(UserID int, req dtos.AddToCartReq, ProductID int) (int, error) {

	cartItem := models.Cart{
		UserID:    UserID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	userID, err := h.productDAO.AddProductToCart(cartItem,ProductID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (h *productService) GetCartItems(UserID int) ([]dtos.CartResp, error) {
	return h.productDAO.GetCartItems(UserID)
}

// func (h *productService) UpdateCartItems(UserID int, ProductID int, req dtos.AddToCartReq) error{

// 	CartItems := models.Cart{
// 		Quantity: req.
// 	}
// }

func (h *productService) AddAddress(Input models.Address, UserID int) (int, error) {
	Address := models.Address{
		UserID: UserID ,
		Name: Input.Name,
		MobileNumber: Input.MobileNumber,
		Street: Input.Street,
		Landmark: Input.Landmark,
		City: Input.City,
		Pincode: Input.Pincode,
		State: Input.State,
		Country: Input.Country,
	}

	return h.productDAO.AddAddress(Address)
}