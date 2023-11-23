package todo

import (
	"context"
	"github.com/linksoft-dev/single/comps/go/dao"
	"github.com/linksoft-dev/single/crud/apps/app"
)

const (
	tableName = "todo"
)

type Repository interface {
	dao.Crud[Model]
}

type Dao struct {
	dao.Crud[Model]
}

func NewRepository(ctx context.Context) (r Repository, err error) {
	// try to get tenantId from context
	tenantId, ok := ctx.Value("tenantId").(string)
	if !ok {
		tenantId = "test"
	}
	r, err = app.GetJsonbDb[Model](ctx, tenantId, tableName)
	return r, nil
}
