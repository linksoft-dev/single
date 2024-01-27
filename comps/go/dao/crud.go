package dao

import (
	"context"
	"fmt"
	"github.com/linksoft-dev/single/comps/go/filter"
	"github.com/linksoft-dev/single/comps/go/types"
	"time"
)

// BaseModel is a model used by most of models having very basic and common Fields for most cases
type BaseModel struct {
	Id             string     `json:"id,omitempty"`
	UserId         string     `json:"userId,omitempty"`   //id do usuario que criou o registro
	UserName       string     `json:"userName,omitempty"` //nome do usuario que criou o registro
	CreatedAt      types.Time `json:"createdAt,omitempty"`
	UserUpdateId   string     `json:"userUpdateId,omitempty"`   //id do usuario que atualizou pela ultima vez
	UserUpdateName string     `json:"userUpdateName,omitempty"` //nome do usuario que atualizou pela ultima vez
	UpdatedAt      types.Time `json:"updatedAt,omitempty"`
	ChangedFields  []string   `json:"-"`
	// soft delete, used to set that some record was deleted, this is used by filter for do not get records
	// if this field set
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

func (b BaseModel) GetId() string {
	return b.Id
}

type ObjI[T any] interface {
	GetId() string
}

var (
	ErrNotFound = fmt.Errorf("record not found")
)

type UpdateField []string

type Crud[T ObjI[T]] interface {
	Create(ctx context.Context, obj T) (T, error)
	Update(ctx context.Context, obj T, fields UpdateField) error
	Save(ctx context.Context, insert bool, obj ...T) ([]T, error)
	Delete(ctx context.Context, id string) error
	DeleteHard(ctx context.Context, id string) error
	List(ctx context.Context, filter filter.Filter) ([]T, error)
	Get(ctx context.Context, id string) (T, error)
}
