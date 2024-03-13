package models

type Paging struct {
	Page    int   `json:"page"`
	PerPage int   `json:"perPage"`
	Counter int64 `json:"counter"`
}

type FilterPagination struct {
	Page     int    `query:"page"`
	PerPage  int    `query:"per_page"`
	OrderBy  string `query:"order_by"`
	SortType string `query:"sort_type"`
}
