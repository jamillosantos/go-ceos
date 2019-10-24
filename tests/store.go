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

type userGroupStore struct {
	*ceous.BaseStore
}

func NewUserGroupStore(options ...ceous.CeousOption) *userGroupStore {
	return &userGroupStore{
		BaseStore: ceous.NewStore(userGroupBaseSchema, options...),
	}
}

func (store *userGroupStore) Insert(record *UserGroup, fields ...ceous.SchemaField) error {
	return store.BaseStore.Insert(record, fields...)
}

func (store *userGroupStore) Update(record *UserGroup, fields ...ceous.SchemaField) (int64, error) {
	return store.BaseStore.Update(record, fields...)
}
