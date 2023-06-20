package models

type PaginationParams struct {
	Page     int    `query:"page"`
	PageSize int    `query:"pageSize"`
	Search   string `query:"search"`
	SortBy   string `query:"sortBy"`
	GetAll   bool   `query:"getAll"`
}

type ResponseMetadata struct {
	Page     int  `json:"page"`
	PageSize int  `json:"pageSize"`
	Length   int  `json:"length"`
	HasNext  bool `json:"hasNext"`
}

type PaginatedResponse[T any] struct {
	Metadata ResponseMetadata `json:"metadata"`
	Data     []T              `json:"data"`
}
