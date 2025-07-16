package dto

import "time"

// SwapRequestCreateRequest represents the data for creating a swap request.
type SwapRequestCreateRequest struct {
	RequestedItemID uint   `json:"requested_item_id" binding:"required"`
	OfferedItems    []uint `json:"offered_items" binding:"required,min=1,max=10"`
}

// SwapRequestCreateResponse is the response after creating a swap request.
type SwapRequestCreateResponse struct {
	Message       string    `json:"message"`
	SwapRequestID uint      `json:"swap_request_id"`
	Status        string    `json:"status"`
	ExpiresAt     time.Time `json:"expires_at"`
}

// SwapRequestResponse defines the structure for a single swap request.
type SwapRequestResponse struct {
	ID            uint           `json:"id"`
	Requester     UserResponse   `json:"requester"`
	Owner         UserResponse   `json:"owner"`
	RequestedItem ItemResponse   `json:"requested_item"`
	OfferedItems  []ItemResponse `json:"offered_items"`
	Status        string         `json:"status"`
	CreatedAt     time.Time      `json:"created_at"`
	ExpiresAt     time.Time      `json:"expires_at"`
}

// SwapRequestsResponse defines the structure for a list of swap requests.
type SwapRequestsResponse struct {
	SwapRequests []SwapRequestResponse `json:"swap_requests"`
}

// SwapRequestUpdateRequest represents the data for updating a swap request.
type SwapRequestUpdateRequest struct {
	Action string `json:"action" binding:"required,oneof=accept,decline"`
}

// SwapRequestUpdateResponse is the response after updating a swap request.
type SwapRequestUpdateResponse struct {
	Message       string `json:"message"`
	SwapRequestID uint   `json:"swap_request_id"`
	Status        string `json:"status"`
}

// SwapCompleteResponse is the response after completing a swap.
type SwapCompleteResponse struct {
	Message           string `json:"message"`
	SwapTransactionID uint   `json:"swap_transaction_id"`
	Status            string `json:"status"`
}
