package todo

import "github.com/linksoft-dev/single/comps/go/dao"

type Model struct {
	dao.BaseModel
	Id          string
	Description string
}
