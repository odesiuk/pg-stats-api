package rest

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	FnGetter func(key string, defaultValue ...string) string

	ParamGetter interface {
		Params(key string, defaultValue ...string) string
	}

	QueryGetter interface {
		Query(key string, defaultValue ...string) string
	}

	Parser struct {
		source FnGetter
	}
)

func Param(s ParamGetter) *Parser {
	return &Parser{source: s.Params}
}

func QueryParam(s QueryGetter) *Parser {
	return &Parser{source: s.Query}
}

func (p *Parser) StringFromEnum(name string, enum []string) (string, error) {
	val := p.source(name)
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

func (p *Parser) String(name string) (string, error) {
	val := p.source(name)
	if val != "" {
		return val, nil
	}

	return "", ErrBadRequestInvalidParameter(name)
}

func (p *Parser) IntOrDefault(name string, def int) (int, error) {
	val := p.source(name)
	if val == "" {
		return def, nil
	}

	intVal, err := strconv.Atoi(val)
	if err == nil {
		return intVal, nil
	}

	return 0, ErrBadRequestInvalidParameter(name).WithErr(err)
}

func (p *Parser) SimplePagination() (SimplePaginationParams, error) {
	var (
		err error
		pag = SimplePaginationParams{Page: DefaultPage}
	)

	page := p.source(PaginationParamPage)
	if page == "" {
		return pag, nil
	}

	pag.Page, err = strconv.Atoi(page)
	if err != nil {
		return pag, ErrBadRequestInvalidParameter(PaginationParamPage).WithErr(err)
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
