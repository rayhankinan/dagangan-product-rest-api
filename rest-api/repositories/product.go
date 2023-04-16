package repositories

type SearchProductsQuery struct {
	Search string `form:"search" field:"search"`
	Sort   string `form:"sort" field:"sort"`
	Page   int    `form:"page" field:"page" binding:"required,gt=0"`
	Size   int    `form:"size" field:"size" binding:"required,gt=0"`
}

type AddProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       uint   `json:"price" binding:"required,gte=0"`
	Stock       uint   `json:"stock" binding:"required,gte=0"`
}

type EditProductRequest struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Price       uint   `json:"price" binding:"gte=0"`
	Stock       uint   `json:"stock" binding:"gte=0"`
}
