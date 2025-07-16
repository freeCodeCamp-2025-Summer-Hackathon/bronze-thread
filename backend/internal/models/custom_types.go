package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Custom types for better validation
type ItemCondition string
type ItemStatus string
type SwapRequestStatus string

const (
	// Item conditions
	ConditionNew     ItemCondition = "new"
	ConditionLikeNew ItemCondition = "like_new"
	ConditionGood    ItemCondition = "good"
	ConditionFair    ItemCondition = "fair"
	ConditionPoor    ItemCondition = "poor"

	// Item statuses
	StatusAvailable ItemStatus = "available"
	StatusPending   ItemStatus = "pending"
	StatusSwapped   ItemStatus = "swapped"

	// Swap request statuses
	SwapPending   SwapRequestStatus = "pending"
	SwapAccepted  SwapRequestStatus = "accepted"
	SwapRejected  SwapRequestStatus = "rejected"
	SwapCancelled SwapRequestStatus = "cancelled"
	SwapCompleted SwapRequestStatus = "completed"
)

// StringArray is a custom type for handling JSON arrays in database
type StringArray []string

func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = StringArray{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, s)
}

func (s StringArray) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "[]", nil
	}
	return json.Marshal(s)
}

// Validation methods for enums
func (ic ItemCondition) IsValid() bool {
	switch ic {
	case ConditionNew, ConditionLikeNew, ConditionGood, ConditionFair, ConditionPoor:
		return true
	}
	return false
}

func (is ItemStatus) IsValid() bool {
	switch is {
	case StatusAvailable, StatusPending, StatusSwapped:
		return true
	}
	return false
}

func (srs SwapRequestStatus) IsValid() bool {
	switch srs {
	case SwapPending, SwapAccepted, SwapRejected, SwapCancelled, SwapCompleted:
		return true
	}
	return false
}
