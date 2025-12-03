package repos

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/terrabase-dev/terrabase/internal/models"
	"github.com/uptrace/bun"
	"google.golang.org/protobuf/proto"
)

func uuidString() string {
	return uuidString()
}

func rowCount(res sql.Result) int64 {
	if res == nil {
		return 0
	}

	n, _ := res.RowsAffected()

	return n
}

func create[T models.TerrabaseModel[P], P proto.Message](ctx context.Context, db *bun.DB, model T) (P, error) {
	_, err := db.NewInsert().Model(model).Returning("*").Exec(ctx)
	if err != nil {
		var zero P

		if isUniqueViolation(err) {
			return zero, fmt.Errorf("%s: %w", model.ModelName(), ErrAlreadyExists)
		}

		return zero, err
	}

	return model.ToProto()
}

func get[T models.TerrabaseModel[P], P proto.Message](ctx context.Context, query *bun.SelectQuery, model T, id string) (P, error) {
	err := query.Where("id = ?", id).Scan(ctx)
	if err != nil {
		var zero P

		if isNotFound(err) {
			return zero, fmt.Errorf("%s: %w", model.ModelName(), ErrNotFound)
		}

		return zero, err
	}

	return model.ToProto()
}

func paginate[T models.TerrabaseModel[P], P proto.Message](ctx context.Context, query *bun.SelectQuery, queryResults *[]T, pageSize int32, pageToken string) ([]P, string, error) {
	limit := int(pageSize)

	if limit <= 0 || limit > 1000 {
		limit = 50
	}

	offset := 0

	if pageToken != "" {
		if parsed, err := strconv.Atoi(pageToken); err == nil && parsed >= 0 {
			offset = parsed
		}
	}

	err := query.Limit(limit).Offset(offset).Scan(ctx)
	if err != nil {
		return nil, "", err
	}

	results := make([]P, 0, len(*queryResults))

	for i := range *queryResults {
		proto, err := (*queryResults)[i].ToProto()
		if err != nil {
			return nil, "", err
		}

		results = append(results, proto)
	}

	nextToken := ""

	if len(*queryResults) == limit {
		nextToken = strconv.Itoa(offset + limit)
	}

	return results, nextToken, nil
}

func update[T models.TerrabaseModel[P], P proto.Message](ctx context.Context, db *bun.DB, model T, columns ...string) (P, error) {
	model.SetUpdatedAt(time.Now().UTC())

	columns = append(columns, "updated_at")

	_, err := db.NewUpdate().Model(model).Column(columns...).WherePK().Exec(ctx)
	if err != nil {
		var zero P

		if isUniqueViolation(err) {
			return zero, fmt.Errorf("%s: %w", model.ModelName(), ErrAlreadyExists)
		}

		return zero, err
	}

	return model.ToProto()
}

func delete[T models.TerrabaseModel[P], P proto.Message](ctx context.Context, db *bun.DB, model T, id string) error {
	res, err := db.NewDelete().Model(model).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	if rowCount(res) == 0 {
		return ErrNotFound
	}

	return nil
}
