package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// ValidateEmail validates email format
func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePhone validates phone number format (simple validation)
func ValidatePhone(phone string) bool {
	// Remove any spaces, dashes, or parentheses
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")
	
	// Check if it contains only digits and has a reasonable length
	re := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	return re.MatchString(phone)
}

// FormatCurrency formats a float64 amount as currency string
func FormatCurrency(amount float64) string {
	return "Rp " + FormatNumber(amount)
}

// FormatNumber formats a float64 as a string with thousand separators
func FormatNumber(number float64) string {
	// For now, we'll just return the raw number as string
	// In a real implementation, you might want to format it with thousand separators
	return FormatFloat(number, 2)
}

// FormatFloat formats a float64 to a specific decimal places
func FormatFloat(f float64, precision int) string {
	// Using Go's fmt package for formatting
	return fmt.Sprintf("%."+fmt.Sprint(precision)+"f", f)
}

// ValidateStruct validates a struct using tags
func ValidateStruct(s interface{}) error {
	// Placeholder implementation - in a real app you might use a validation library like validator
	return nil
}

// Contains checks if a string is in a slice of strings
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// IsValidTransactionCategory checks if a category is valid
func IsValidTransactionCategory(category string) bool {
	validCategories := []string{"RENT", "SALARY", "EQUIPMENT", "PAYMENT", "OTHER"}  // These values must match the constant values in models
	return Contains(validCategories, category)
}

// IsValidOrderStatus checks if status is valid
func IsValidOrderStatus(status string) bool {
	validStatuses := []string{"pending", "confirmed", "cancelled", "completed"}
	return Contains(validStatuses, status)
}

// IsValidPaymentStatus checks if payment status is valid
func IsValidPaymentStatus(status string) bool {
	validStatuses := []string{"unpaid", "partially_paid", "paid", "overdue", "cancelled"}
	return Contains(validStatuses, status)
}

// IsValidResellerStatus checks if reseller status is valid
func IsValidResellerStatus(status string) bool {
	validStatuses := []string{"active", "inactive"}
	return Contains(validStatuses, status)
}

// IsValidProductStatus checks if product status is valid
func IsValidProductStatus(status string) bool {
	validStatuses := []string{"active", "inactive"}
	return Contains(validStatuses, status)
}

// ValidateUUID checks if a string is a valid UUID
func ValidateUUID(uuid string) bool {
	// Basic UUID validation - 8-4-4-4-12 hex characters
	uuidRegex := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	return uuidRegex.MatchString(uuid)
}

// Error helpers
var (
	ErrInvalidEmail          = errors.New("invalid email format")
	ErrInvalidPhone          = errors.New("invalid phone format")
	ErrInvalidAmount         = errors.New("amount must be greater than zero")
	ErrInsufficientStock     = errors.New("insufficient stock")
	ErrOrderNotFound         = errors.New("order not found")
	ErrProductNotFound       = errors.New("product not found")
	ErrResellerNotFound      = errors.New("reseller not found")
	ErrPaymentNotFound       = errors.New("payment not found")
	ErrInvalidOrderStatus    = errors.New("invalid order status")
	ErrInvalidPaymentStatus  = errors.New("invalid payment status")
	ErrInvalidTransactionCategory = errors.New("invalid transaction category")
)