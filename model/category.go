package model

// Category represents a category entity.
type Category struct {
	ID          int64   `json:"id"`
	UUID        string  `json:"uuid"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
