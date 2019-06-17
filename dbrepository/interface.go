package dbrepository

import "MAD_Rest_Assign/domain"

type Reader interface {
	GetByID(id domain.ID) (*domain.Restaurant, error)
	GetAll() ([]*domain.Restaurant, error)
	FindByName(name string) ([]*domain.Restaurant, error)
}

type Writer interface {
	Update(inp *domain.Restaurant) error
	Store(b *domain.Restaurant) (domain.ID, error)
	Delete(id domain.ID) error
}

type Filter interface {
	FindByTypeOfFood(foodType string) ([]*domain.Restaurant, error)
	FindByTypeOfPostCode(postCode string) ([]*domain.Restaurant, error)
	Search(query string) ([]*domain.Restaurant, error)
}

type Repository interface {
	Reader
	Writer
	Filter
}
