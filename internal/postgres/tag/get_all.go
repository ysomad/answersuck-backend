package tag

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/ysomad/answersuck/internal/entity"
	"github.com/ysomad/answersuck/internal/pkg/paging"
	"github.com/ysomad/answersuck/internal/pkg/pgsearch"
	"github.com/ysomad/answersuck/internal/pkg/sort"
)

func (r *repository) GetAll(ctx context.Context, search string, p paging.OffsetParams, sorts []sort.Sort) (paging.List[entity.Tag], error) {
	b := r.Builder.
		Select("name, author, created_at").
		From(tagTable).
		Limit(p.Limit + 1).
		Offset(p.Offset)

	if search != "" {
		b = pgsearch.Where(b, search, pgsearch.ConfRus)
		b = pgsearch.OrderBy(b, search, pgsearch.ConfRus, pgsearch.OrderDefault)
	}

	for _, sort := range sorts {
		b = sort.Attach(b)
	}

	sql, args, err := b.ToSql()
	if err != nil {
		return paging.List[entity.Tag]{}, fmt.Errorf("b.ToSql: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return paging.List[entity.Tag]{}, fmt.Errorf("r.Pool.Query: %w", err)
	}

	tags, err := pgx.CollectRows(rows, pgx.RowToStructByPos[entity.Tag])
	if err != nil {
		return paging.List[entity.Tag]{}, fmt.Errorf("pgx.RowToStructPyBos: %w", err)
	}

	return paging.NewList(tags, p.Limit)
}