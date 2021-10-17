package rest

const (
	PaginationParamPage  = "page"
	PaginationParamLimit = "limit"
	DefaultLimit         = 10
	DefaultPage          = 1
)

type SimplePaginationParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
