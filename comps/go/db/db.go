package db

import (
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

type UpdateField []string

type Filter struct {
}

type Crud[T ObjI[T]] interface {
	Create(obj T) (T, error)
	Update(obj T, fields UpdateField) error
	Save(obj ...T) ([]T, error)
	Delete(T) error
	DeleteHard(T) error
	Find(filter Query) ([]T, error)
	Get(id string) (T, error)
}
