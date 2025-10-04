package service

import (
	"oracle-crud-api/internal/model"
	repository "oracle-crud-api/internal/repo"
)

type PersonService struct {
	repo *repository.PersonRepository
}

func NewPersonService(repo *repository.PersonRepository) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) Create(p *model.Person) error {
	return s.repo.Create(p)
}

func (s *PersonService) GetByID(id string) (*model.Person, error) {
	return s.repo.GetByID(id)
}

func (s *PersonService) GetAll() ([]model.Person, error) {
	return s.repo.GetAll()
}

func (s *PersonService) Update(p *model.Person) error {
	return s.repo.Update(p)
}

func (s *PersonService) Delete(id string) error {
	return s.repo.Delete(id)
}
