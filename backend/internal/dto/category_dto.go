package dto

// CategoryResponse defines the structure for a single category.
type CategoryResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

// CategoriesResponse defines the structure for a list of categories.
type CategoriesResponse struct {
	Categories []CategoryResponse `json:"categories"`
}
