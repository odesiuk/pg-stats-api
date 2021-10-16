package rest

import (
	"strconv"
	"time"
)

const (
	dateFormat = "2006-01-02"
	uintBase   = 10
)

type (
	FnGetter func(string) string

	ParamGetter interface {
		Param(string) string
	}

	QueryParamGetter interface {
		QueryParam(string) string
	}

	Parser struct {
		source FnGetter
	}
)

func Param(s ParamGetter) *Parser {
	return &Parser{source: s.Param}
}

func QueryParam(s QueryParamGetter) *Parser {
	return &Parser{source: s.QueryParam}
}

func (p *Parser) StringFromEnum(name string, enum []string) (string, error) {
	val := p.source(name)
	for _, s := range enum {
		if val == s {
			return s, nil
		}
	}

	return "", ErrBadRequestInvalidParameter(name)
}

func (p *Parser) String(name string) (string, error) {
	val := p.source(name)
	if val != "" {
		return val, nil
	}

	return "", ErrBadRequestInvalidParameter(name)
}

func (p *Parser) StringOrDefault(name, def string) string {
	if val := p.source(name); val != "" {
		return val
	}

	return def
}

func (p *Parser) Int(name string) (int, error) {
	intVal, err := strconv.Atoi(p.source(name))
	if err == nil {
		return intVal, nil
	}

	return 0, ErrBadRequestInvalidParameter(name)
}

func (p *Parser) Uint64(name string) (uint64, error) {
	n, err := strconv.ParseUint(p.source(name), uintBase, 64)
	if err != nil {
		return 0, ErrBadRequestInvalidParameter(name).WithErr(err)
	}

	return n, nil
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

func (p *Parser) DateOrDefault(name string, def time.Time) (time.Time, error) {
	val := p.source(name)
	if val == "" {
		return def, nil
	}

	dateVal, err := time.Parse(dateFormat, val)
	if err == nil {
		return dateVal, nil
	}

	return time.Time{}, ErrBadRequestInvalidParameter(name).WithErr(err)
}
