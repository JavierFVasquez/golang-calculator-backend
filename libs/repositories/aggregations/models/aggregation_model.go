package aggregationModels

type AggregationPaginationOptions struct {
	Skip     int
	Limit    int
	Sort     string
	GetAll   bool
	Page     int
	PageSize int
	Search   string
}
