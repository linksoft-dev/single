package db

type Crud interface {
	Create(obj interface{}) error
	Update(obj interface{}) error
}
