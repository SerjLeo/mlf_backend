package models

import (
	"fmt"
	"strings"
)

type PaginationParams struct {
	Page    int `json:"page" form:"page"`
	PerPage int `json:"perPage" form:"perPage"`
}

func (p *PaginationParams) ToString() string {
	result := ""

	query := make([]string, 0, 2)

	if p.Page != 0 {
		query = append(query, fmt.Sprintf("page=%d", p.Page))
	}

	if p.PerPage != 0 {
		query = append(query, fmt.Sprintf("perPage=%d", p.PerPage))
	}

	result += strings.Join(query, "&")

	return result
}
