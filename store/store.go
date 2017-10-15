package store

import (
	"github.com/jmoiron/sqlx"

	"github.com/steven-ferrer/go-db-models/model"
)

//UserStore implements models.UserStorer
type UserStore struct {
	db *sqlx.DB
}

func (cs *UserStore) List(filters ...models.UserFilter) ([]models.User, error) {
	return nil, nil
}

func (cs *UserStore) Get(filters ...models.UserFilter) (models.User, error) {
	return models.User{}, nil
}

func (cs *UserStore) Create(User models.User) error {
	return nil
}

func (cs *UserStore) Update(User models.User) error {
	return nil
}

func (cs *UserStore) Delete(filters ...models.UserFilter) error {
	return nil
}

func NewUserStore(db *sqlx.DB) models.UserStorer {
	return &UserStore{db}
}
