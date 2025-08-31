package dtos

type RegisterUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AddToCartReq struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CartResp struct {
	ProductId   int     `json:"product_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
}
