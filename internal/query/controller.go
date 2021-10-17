package query

import (
	"github.com/gofiber/fiber/v2"
	"github.com/odesiuk/pg-stats-api/internal/query/dto"
	"github.com/odesiuk/pg-stats-api/internal/storage/models"
	repos "github.com/odesiuk/pg-stats-api/internal/storage/repositories"
	"github.com/odesiuk/pg-stats-api/pkg/rest"
)

const (
	FilterParamType = "type"
	SortParam       = "time_spent_sort"
)

var (
	sortEnum           = []string{repos.SortAsc, repos.SortDesc}
	statementTypesEnum = []string{
		models.StatementINSERT,
		models.StatementUPDATE,
		models.StatementSELECT,
		models.StatementDELETE,
	}
)

type (
	StatementsRepo interface {
		GetByType(minExecTime uint64, qType string, timeSpentSort string, limit, offset int) ([]models.PgStatStatement, error)
	}

	Controller struct {
		repo             StatementsRepo
		minQueryDuration uint64
	}
)

func NewController(repo StatementsRepo, minQueryDuration uint64) Controller {
	return Controller{repo, minQueryDuration}
}

func (c Controller) GetAll(ctx *fiber.Ctx) error {
	p := rest.QueryParam(ctx)

	// parse statement type param.
	qType, err := p.StringFromEnum(FilterParamType, statementTypesEnum)
	if err != nil {
		return err
	}

	// parse pagination params.
	pagination, err := p.SimplePagination()
	if err != nil {
		return err
	}

	// parse time_spent sort param.
	sort, err := p.StringFromEnum(SortParam, sortEnum)
	if err != nil {
		return err
	}

	// get from storage.
	statements, err := c.repo.GetByType(
		c.minQueryDuration,
		qType, sort,
		pagination.Limit,
		pagination.Limit*(pagination.Page-1),
	)
	if err != nil {
		return err
	}

	// convert to DTO
	queries := make([]dto.Query, len(statements))
	for i, statement := range statements {
		queries[i] = dto.Query{
			ID:           statement.QueryID,
			Statement:    statement.Query,
			MaxExecTime:  statement.MaxExecTime,
			MeanExecTime: statement.MeanExecTime,
		}
	}

	return ctx.JSON(rest.ResponseWithPayload{
		Pagination: &rest.Pagination{
			Page:     pagination.Page,
			PageSize: len(queries),
		},
		Payload: queries,
	})
}
