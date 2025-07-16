package controllers

import (
	"net/http"
	"strconv"

	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/database"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/dto"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/models"
	"github.com/gin-gonic/gin"
)

// CreateItem handles the creation of a new item listing.
func CreateItem(c *gin.Context) {
	// Get user ID from the JWT token (set by the IsAuthorized middleware)
	userID, exists := c.Get("sub")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Unauthorized"})
		return
	}

	var req dto.ItemCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Validation failed", Message: err.Error()})
		return
	}

	// Create a new item model
	item := models.Item{
		OwnerID:        userID.(uint),
		CategoryID:     req.CategoryID,
		Name:           req.Name,
		Description:    req.Description,
		Condition:      models.ItemCondition(req.Condition),
		EstimatedValue: req.EstimatedValue,
		Currency:       req.Currency,
		ImageURLs:      req.ImageURLs,
		Status:         models.StatusAvailable, // Default status
	}

	// Save the new item to the database
	if err := database.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create item"})
		return
	}

	c.JSON(http.StatusCreated, dto.ItemCreateResponse{Message: "Item listed successfully", ItemID: item.ID})
}

// GetAllItems handles retrieving all available items with optional filters.
func GetAllItems(c *gin.Context) {
	var params dto.ItemQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid query parameters", Message: err.Error()})
		return
	}
	params.SetDefaults() // Apply default pagination if not provided

	query := database.DB.Model(&models.Item{}).Preload("Owner").Preload("Category")

	// Apply filters based on query parameters
	if params.CategoryID != nil {
		query = query.Where("category_id = ?", *params.CategoryID)
	}
	if params.Search != nil {
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+*params.Search+"%", "%"+*params.Search+"%")
	}
	if params.Condition != nil {
		query = query.Where("condition = ?", *params.Condition)
	}
	if params.Status != nil {
		query = query.Where("status = ?", *params.Status)
	} else {
		// Default to showing only available items if status is not specified
		query = query.Where("status = ?", models.StatusAvailable)
	}

	// Get total count for pagination
	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to count items"})
		return
	}

	// Apply pagination
	var items []models.Item
	offset := (params.Page - 1) * params.Limit
	if err := query.Offset(offset).Limit(params.Limit).Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to retrieve items"})
		return
	}

	// Map models to DTOs
	var itemResponses []dto.ItemResponse
	for _, item := range items {
		itemResponses = append(itemResponses, dto.ItemResponse{
			ID:             item.ID,
			Name:           item.Name,
			Description:    item.Description,
			Condition:      string(item.Condition),
			EstimatedValue: item.EstimatedValue,
			Currency:       item.Currency,
			ImageURLs:      item.ImageURLs,
			Status:         string(item.Status),
			CreatedAt:      item.CreatedAt,
			UpdatedAt:      item.UpdatedAt,
			Owner: dto.UserResponse{
				ID:        item.Owner.ID,
				Name:      item.Owner.Name,
				AvatarURL: item.Owner.AvatarURL,
				RatingAvg: item.Owner.RatingAvg,
			},
			Category: dto.CategoryResponse{
				ID:   item.Category.ID,
				Name: item.Category.Name,
			},
		})
	}

	c.JSON(http.StatusOK, dto.ItemsResponse{
		Items:        itemResponses,
		TotalResults: int(total),
	})
}

// GetItemByID handles retrieving a single item by its ID.
func GetItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid item ID"})
		return
	}

	var item models.Item
	// Preload owner and category for a full response
	if err := database.DB.Preload("Owner").Preload("Category").First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Item not found"})
		return
	}

	// Map to DTO
	response := dto.ItemResponse{
		ID:             item.ID,
		Name:           item.Name,
		Description:    item.Description,
		Condition:      string(item.Condition),
		EstimatedValue: item.EstimatedValue,
		Currency:       item.Currency,
		ImageURLs:      item.ImageURLs,
		Status:         string(item.Status),
		CreatedAt:      item.CreatedAt,
		UpdatedAt:      item.UpdatedAt,
		Owner: dto.UserResponse{
			ID:        item.Owner.ID,
			Name:      item.Owner.Name,
			AvatarURL: item.Owner.AvatarURL,
			RatingAvg: item.Owner.RatingAvg,
		},
		Category: dto.CategoryResponse{
			ID:   item.Category.ID,
			Name: item.Category.Name,
		},
	}

	c.JSON(http.StatusOK, response)
}

// UpdateItem handles updating an existing item.
func UpdateItem(c *gin.Context) {
	userID, exists := c.Get("sub")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Unauthorized"})
		return
	}

	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid item ID"})
		return
	}

	var req dto.ItemUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Validation failed", Message: err.Error()})
		return
	}

	var item models.Item
	if err := database.DB.First(&item, itemID).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Item not found"})
		return
	}

	// Security Check: Ensure the user owns the item
	if item.OwnerID != userID.(uint) {
		c.JSON(http.StatusForbidden, dto.ErrorResponse{Error: "Forbidden", Message: "You do not have permission to update this item"})
		return
	}

	// Update fields from request
	if req.Name != nil {
		item.Name = *req.Name
	}
	if req.Description != nil {
		item.Description = *req.Description
	}
	if req.Condition != nil {
		item.Condition = models.ItemCondition(*req.Condition)
	}
	// ... add other updatable fields as needed

	if err := database.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update item"})
		return
	}

	c.JSON(http.StatusOK, dto.ItemUpdateResponse{Message: "Item updated successfully", ItemID: item.ID})
}

// DeleteItem handles deleting an item.
func DeleteItem(c *gin.Context) {
	userID, exists := c.Get("sub")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Unauthorized"})
		return
	}

	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid item ID"})
		return
	}

	var item models.Item
	if err := database.DB.First(&item, itemID).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Item not found"})
		return
	}

	// Security Check: Ensure the user owns the item
	if item.OwnerID != userID.(uint) {
		c.JSON(http.StatusForbidden, dto.ErrorResponse{Error: "Forbidden", Message: "You do not have permission to delete this item"})
		return
	}

	if err := database.DB.Delete(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to delete item"})
		return
	}

	c.JSON(http.StatusOK, dto.ItemDeleteResponse{Message: "Item deleted successfully"})
}
