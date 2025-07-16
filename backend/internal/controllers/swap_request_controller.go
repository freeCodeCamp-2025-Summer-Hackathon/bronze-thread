package controllers

import (
	"net/http"
	"strconv"

	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/database"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/dto"
	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateSwapRequest handles initiating a new swap request.
func CreateSwapRequest(c *gin.Context) {
	requesterID, exists := c.Get("sub")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Unauthorized"})
		return
	}

	var req dto.SwapRequestCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Validation failed", Message: err.Error()})
		return
	}

	var requestedItem models.Item
	if err := database.DB.First(&requestedItem, req.RequestedItemID).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Requested item not found"})
		return
	}

	// Business logic checks
	if requestedItem.OwnerID == requesterID.(uint) {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "You cannot swap for your own item"})
		return
	}
	if requestedItem.Status != models.StatusAvailable {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Item is not available for swapping"})
		return
	}

	// Use a transaction to ensure all or nothing is saved
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// Create the swap request
		swapRequest := models.SwapRequest{
			RequesterID:     requesterID.(uint),
			OwnerID:         requestedItem.OwnerID,
			RequestedItemID: req.RequestedItemID,
			Status:          models.SwapPending,
		}
		if err := tx.Create(&swapRequest).Error; err != nil {
			return err
		}

		// Add the offered items to the request
		for _, offeredItemID := range req.OfferedItems {
			swapRequestItem := models.SwapRequestItem{
				SwapRequestID: swapRequest.ID,
				ItemID:        offeredItemID,
				ItemType:      "offered",
			}
			if err := tx.Create(&swapRequestItem).Error; err != nil {
				return err
			}
		}

		// Set the response for the outer function
		c.JSON(http.StatusCreated, dto.SwapRequestCreateResponse{
			Message:       "Swap request initiated successfully",
			SwapRequestID: swapRequest.ID,
			Status:        string(swapRequest.Status),
			ExpiresAt:     swapRequest.ExpiresAt,
		})

		return nil // Commit the transaction
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to create swap request", Message: err.Error()})
		return
	}
}

// GetSwapRequests retrieves a list of swap requests for the authenticated user.
func GetSwapRequests(c *gin.Context) {
	userID, exists := c.Get("sub")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Unauthorized"})
		return
	}

	var params dto.SwapRequestQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid query parameters", Message: err.Error()})
		return
	}

	query := database.DB.Preload("Requester").Preload("Owner").Preload("RequestedItem").Preload("SwapRequestItems.Item")

	switch *params.Type {
	case "sent":
		query = query.Where("requester_id = ?", userID)
	case "received":
		query = query.Where("owner_id = ?", userID)
	default:
		// Both sent and received
		query = query.Where("requester_id = ? OR owner_id = ?", userID, userID)
	}

	if params.Status != nil {
		query = query.Where("status = ?", *params.Status)
	}

	var swapRequests []models.SwapRequest
	if err := query.Find(&swapRequests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to retrieve swap requests"})
		return
	}

	// Map to DTOs
	var responses []dto.SwapRequestResponse
	for _, sr := range swapRequests {
		var offeredItems []dto.ItemResponse
		for _, sri := range sr.SwapRequestItems {
			offeredItems = append(offeredItems, dto.ItemResponse{
				ID:   sri.Item.ID,
				Name: sri.Item.Name,
				// ... map other item fields as needed
			})
		}

		responses = append(responses, dto.SwapRequestResponse{
			ID:            sr.ID,
			Status:        string(sr.Status),
			CreatedAt:     sr.CreatedAt,
			ExpiresAt:     sr.ExpiresAt,
			Requester:     dto.UserResponse{ID: sr.Requester.ID, Name: sr.Requester.Name},
			Owner:         dto.UserResponse{ID: sr.Owner.ID, Name: sr.Owner.Name},
			RequestedItem: dto.ItemResponse{ID: sr.RequestedItem.ID, Name: sr.RequestedItem.Name},
			OfferedItems:  offeredItems,
		})
	}

	c.JSON(http.StatusOK, dto.SwapRequestsResponse{SwapRequests: responses})
}

// UpdateSwapRequest handles accepting or declining a swap request.
func UpdateSwapRequest(c *gin.Context) {
	ownerID, exists := c.Get("sub")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Unauthorized"})
		return
	}

	swapRequestID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid swap request ID"})
		return
	}

	var req dto.SwapRequestUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Validation failed", Message: err.Error()})
		return
	}

	var swapRequest models.SwapRequest
	if err := database.DB.First(&swapRequest, swapRequestID).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Swap request not found"})
		return
	}

	// Security Check: Only the owner of the requested item can update the request.
	if swapRequest.OwnerID != ownerID.(uint) {
		c.JSON(http.StatusForbidden, dto.ErrorResponse{Error: "Forbidden", Message: "You do not have permission to update this request"})
		return
	}

	// Update status based on action
	switch req.Action {
	case "accept":
		swapRequest.Status = models.SwapAccepted
	case "decline":
		swapRequest.Status = models.SwapRejected
	default:
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid action"})
		return
	}

	if err := database.DB.Save(&swapRequest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to update swap request"})
		return
	}

	c.JSON(http.StatusOK, dto.SwapRequestUpdateResponse{
		Message:       "Swap request updated successfully",
		SwapRequestID: swapRequest.ID,
		Status:        string(swapRequest.Status),
	})
}
