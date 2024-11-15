package service

import "sync_btim/entity"

type IService interface {
	UpdateClientID(email string, limit int, offset int) (bool, error)
}

type service struct {
	repository entity.IRepository
}

func NewUserService(repository entity.IRepository) *service {
	return &service{repository}
}

func (s *service) UpdateClientID(email string, limit int, offset int) (bool, error) {
	data, err := s.repository.SyncClientID(email, offset, limit)
	if err != nil {
		return data, err
	}

	return data, nil
}
