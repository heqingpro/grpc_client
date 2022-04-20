package main

import (
	"helloword/mock"
)

type JobService struct {
	workService mock.WorkService
}

func NewJobService(service mock.WorkService) *JobService {
	return &JobService{
		workService: service,
	}
}

//
func (s *JobService) Run(name string) string{

	job := s.workService.WorkStart(name)
	
	return job.Name
}

