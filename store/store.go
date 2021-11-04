package store

import (
	"github.com/jmoiron/sqlx"

	"github.com/stevenferrer/go-db-models/model"
)

//UserStore implements models.UserStorer
type UserStore struct {
	db *sqlx.DB
}

func (cs *UserStore) List(filters ...model.UserFilter) ([]model.User, error) {
	return nil, nil
}

func (cs *UserStore) Get(filters ...model.UserFilter) (model.User, error) {
	return model.User{}, nil
}

func (cs *UserStore) Create(User model.User) error {
	return nil
}

func (cs *UserStore) Update(User model.User) error {
	return nil
}

func (cs *UserStore) Delete(filters ...model.UserFilter) error {
	return nil
}

func NewUserStore(db *sqlx.DB) model.UserStorer {
	return &UserStore{db}
}
