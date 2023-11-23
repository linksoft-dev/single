package todo

import (
	"context"
	"github.com/linksoft-dev/single/comps/go/str"
)

type Service struct {
	ctx        context.Context
	repository Repository
}

func NewService(ctx context.Context) (Service, error) {
	var err error
	s := Service{ctx: ctx}
	s.repository, err = NewRepository(ctx)
	if err != nil {
		return Service{}, err
	}
	return s, nil
}

func (s *Service) Create(m *Model) (err error) {
	m.BaseModel.Id = str.Uuid()
	created, err := s.repository.Create(*m)
	if err != nil {
		return err
	}
	*m = created
	return err
}
