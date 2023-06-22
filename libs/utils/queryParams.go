package utils

import (
	"strconv"

	"github.com/JavierFVasquez/truenorth-calculator-backend/libs/models"
)

func ParsePaginationParams(queryParams map[string]string) (*models.PaginationParams) {
	page, pageErr := strconv.Atoi(queryParams["page"])
	if pageErr != nil {
		page = 1
	}
	pageSize, pageSizeErr := strconv.Atoi(queryParams["pageSize"])
	if pageSizeErr != nil {
		pageSize = 20
	}
	getAll, getAllErr := strconv.ParseBool(queryParams["getAll"])
	if getAllErr != nil {
		getAll = false
	}

	paginationOptions := models.PaginationParams{
		Page:     page,
		PageSize: pageSize,
		Search:   queryParams["search"],
		GetAll:   getAll,
		SortBy:   queryParams["sortBy"],
	}

	return &paginationOptions
}
