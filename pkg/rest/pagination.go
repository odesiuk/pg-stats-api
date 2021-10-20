package rest

const (
	PaginationParamPage  = "page"
	PaginationParamLimit = "limit"
	DefaultPage          = 1
)

type SimplePaginationParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
