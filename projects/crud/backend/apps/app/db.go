package app

import (
	"context"
	"fmt"
	"github.com/linksoft-dev/single/comps/go/dao"
	"github.com/linksoft-dev/single/comps/go/dao/postgres/jsonb"
)

// ConnectDatabase this function is an example how to connect to database
func ConnectDatabase() {
	jsonb.Connect(Config.Db.Host,
		Config.Db.Port,
		Config.Db.User,
		Config.Db.Password,
		Config.Db.DbName,
		false,
	)
}

// GetJsonbDb this function returns a jsonb implementation of dao interface dao.Crud
// this jsonb implementation, works with multi-tenant approach
// tenantId will a table name
// resourceName is will be stored in collection column
func GetJsonbDb[T dao.ObjI[T]](ctx context.Context, tenantId, resourceName string) (dao.Crud[T], error) {
	tenantId = fmt.Sprintf("tenant_%s", tenantId)
	db, err := jsonb.NewDataBase[T](ctx, Config.Db.DbName, tenantId, resourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}
