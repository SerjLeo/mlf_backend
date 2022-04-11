package models

type PaginationParams struct {
	Page    int `json:"page" form:"page"`
	PerPage int `json:"perPage" form:"perPage"`
}
