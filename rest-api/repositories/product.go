package repositories

type SearchProductsQuery struct {
	Search string `form:"search" field:"search" binding:"ascii"`
	Page   int    `form:"page" field:"page" binding:"required,gt=0"`
	Size   int    `form:"size" field:"size" binding:"required,gt=0"`
}

type AddProductRequest struct {
	Name        string `json:"name" binding:"required,ascii"`
	Image       string `json:"image" binding:"required,url"`
	Description string `json:"description" binding:"required,ascii"`
	Price       uint   `json:"price" binding:"required,gte=0"`
	Stock       uint   `json:"stock" binding:"required,gte=0"`
}

type EditProductRequest struct {
	Name        string `json:"name" binding:"ascii"`
	Image       string `json:"image" binding:"url"`
	Description string `json:"description" binding:"ascii"`
	Price       uint   `json:"price" binding:"gte=0"`
	Stock       uint   `json:"stock" binding:"gte=0"`
}
