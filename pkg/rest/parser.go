package rest

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	FnGetter func(key string, defaultValue ...string) string

	QueryGetter interface {
		Query(key string, defaultValue ...string) string
	}

	Parser struct {
		source FnGetter
	}
)

func QueryParam(s QueryGetter) *Parser {
	return &Parser{source: s.Query}
}

func (p *Parser) StringFromEnum(name string, enum []string, defaultVal ...string) (string, error) {
	val := p.source(name)
	if val == "" && len(defaultVal) > 0 {
		return defaultVal[0], nil
	}

	for _, s := range enum {
		if val == s {
			return s, nil
		}
	}

	return "", ErrBadRequestInvalidParameter(name, fmt.Sprintf(
		"Available values: %s",
		strings.Join(enum, ", "),
	))
}

func (p *Parser) SimplePagination() (SimplePaginationParams, error) {
	var (
		err error
		pag = SimplePaginationParams{Page: DefaultPage}
	)

	if page := p.source(PaginationParamPage); page != "" {
		pag.Page, err = strconv.Atoi(page)
		if err != nil {
			return pag, ErrBadRequestInvalidParameter(PaginationParamPage).WithErr(err)
		}
	}

	pag.Limit = DefaultLimit
	if limit := p.source(PaginationParamLimit); limit != "" {
		pag.Limit, err = strconv.Atoi(limit)
		if err != nil {
			return pag, ErrBadRequestInvalidParameter(PaginationParamLimit).WithErr(err)
		}
	}

	return pag, nil
}
