package store

import (
	"sync"

	"gorm.io/gorm"
)

var (
	once *sync.Once

	S *datastore
)

type IsStore interface {
	DB() *gorm.DB
	Users() UserStore
	Domains() DomainTaskStore
}

type datastore struct {
	db *gorm.DB
}

var _ IsStore = (*datastore)(nil)

func NewStore(db *gorm.DB) *datastore {
	once.Do(func() {
		S = &datastore{db}
	})
	return S
}

func (ds *datastore) DB() *gorm.DB {
	return ds.db
}

func (ds *datastore) Users() UserStore {
	return newUser(ds.db)
}

func (ds *datastore) Domains() DomainTaskStore {
	return newTasks(ds.db)
}
