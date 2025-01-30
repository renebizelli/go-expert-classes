package main

type Service struct {
	Repo IGetRepo
}

func NewService(repo IGetRepo) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) Get(id int) string {
	return s.Repo.Get(id)
}
