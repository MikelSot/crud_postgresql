package product

import "time"

// model of product
type Model struct {
	Id uint
	Name string
	Observaions string
	Price float32
	CreateAt time.Time
	UpdateAt time.Time
}

// Models slice of Model
type Models []*Model


//	Creamos solo el CRUD en el modelo PRODUCT ya que en esta ocacion es solo en este que se hara.
type Storage interface {
	Create(*Model) error
	UPdate(*Model) error
	GetALl() (Models, error)  // nos va a devolver un SLICE de Models o u error
	GetById(uint) (*Model, error)
	Delete(uint) error
}
