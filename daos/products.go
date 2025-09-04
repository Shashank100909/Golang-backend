package dao

import (
	"github.com/Shashank100909/STUDENTS-API/configs"
	"github.com/Shashank100909/STUDENTS-API/dtos"
	"github.com/Shashank100909/STUDENTS-API/models"
	"gorm.io/gorm"
)

type ProductDAO interface {
	AddProduct(product models.Products) error
	GetProducts(ProductID int) ([]models.Products, error)
	DeleteProduct(ProductID int) error

	AddProductToCart(cartItem models.Cart, ProductID int) (int, error)
	GetCartItems(UserID int) ([]dtos.CartResp, error)
	DeleteProductFromCart(ProductID int, UserID int) error

	AddAddress(Address models.Address) (int, error)
}

type productDao struct {
	db *gorm.DB
}

func NewProductDAO() ProductDAO {
	return &productDao{
		db: configs.DB,
	}
}

func (h *productDao) AddProduct(product models.Products) error {
	return h.db.Create(&product).Error
}

func (h *productDao) GetProducts(ProductID int) ([]models.Products, error) {
	var products []models.Products

	query := h.db.Model(&models.Products{}).
		Select(`
		 id,
		 name,
		 price,
		 description
		 `)

	if ProductID != 0 {
		query = query.Where("id = ?", ProductID)
	}

	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return products, nil
}

func (h *productDao) DeleteProduct(ProductID int) error {

	tx := h.db.Begin()

	err := h.db.Where("id = ?", ProductID).First(&models.Products{}).Error
	if err != nil {
		return gorm.ErrRecordNotFound
	}

	if err = h.db.Where("product_id = ?", ProductID).Delete(&models.Cart{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err = h.db.Where("id = ?", ProductID).Delete(&models.Products{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (h *productDao) AddProductToCart(cartItem models.Cart, ProductID int) (int, error) {
	var count int64

	err := h.db.Model(&models.Cart{}).
		Where("user_id = ? AND product_id = ?", cartItem.UserID, cartItem.ProductID).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	if count > 0 {
		err = h.db.Model(&models.Cart{}).
			Where("user_id = ? AND product_id = ?", cartItem.UserID, cartItem.ProductID).
			Update("quantity", cartItem.Quantity).Error
		if err != nil {
			return 0, err
		}
	} else {
		err = h.db.Create(&cartItem).Error
		if err != nil {
			return 0, err
		}
	}

	return cartItem.UserID, nil
}

func (h *productDao) GetCartItems(UserID int) ([]dtos.CartResp, error) {
	var CartItems []dtos.CartResp

	query := h.db.Model(models.Cart{}).
		Select(`
		carts.product_id,
		products.name,
		products.price,
		products.description,
		carts.quantity
	`).Joins("JOIN products ON carts.product_id = products.id")

	err := query.Where("user_id = ?", UserID).Find(&CartItems).Error
	if err != nil {
		return nil, err
	}

	return CartItems, nil
}

func (h *productDao) DeleteProductFromCart(ProductID int, UserID int) error {
	result := h.db.Where("product_id = ? AND user_id = ?", ProductID, UserID).
		Delete(&models.Cart{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (h *productDao) AddAddress(Address models.Address) (int, error) {
	err := h.db.Create(&Address).Error
	if err != nil {
		return 0, err
	}
	return Address.AddressID, nil
}
