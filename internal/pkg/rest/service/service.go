package service

import (
	. "fio"
	"fio/internal/pkg/rest/repository"
)

type Record interface {
	CreateClient(client Client) (int, error)
	GetClientsByFilter(filter ClientFilter, page int) ([]Client, error)
	UpdateClientRecord(id int, client ClientUpdate) error
	DeleteClientById(id int) error
}

type Service struct {
	Record
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Record: NewRecordService(repos.Record),
	}
}
