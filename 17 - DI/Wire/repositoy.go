package main

import "strconv"

type ServiceRepository struct{}

type IGetRepo interface {
	Get(id int) string
}

func NewServiceRepo() *ServiceRepository {
	return &ServiceRepository{}
}

func (s *ServiceRepository) Get(id int) string { return strconv.Itoa(id) }
