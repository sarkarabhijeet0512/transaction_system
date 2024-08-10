package model

type (
	GenericRes struct {
		Success bool       `json:"success"`
		Message string     `json:"message"`
		Data    any        `json:"data,omitempty"`
		Meta    Pagination `json:"meta,omitempty"`
	}
	Pagination struct {
		CurrentPage    int `json:"current_page,omitempty"`
		TotalPages     int `json:"total_pages,omitempty"`
		TotalDataCount int `json:"total_data_count,omitempty"`
	}
)
