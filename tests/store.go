package tests

import "github.com/jamillosantos/go-ceous"

type userStore struct {
	*ceous.BaseStore
}

func NewUserStore(options ...ceous.CeousOption) *userStore {
	return &userStore{
		BaseStore: ceous.NewStore(userBaseSchema, options...),
	}
}

func (store *userStore) Insert(record *User, fields ...ceous.SchemaField) error {
	return store.BaseStore.Insert(record, fields...)
}
