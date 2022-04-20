package mock

import "fmt"

type WorkService interface{
	WorkStart(name string) *Job
}

type Work struct {
	WorkName  string
}

type Job struct {
	Name string
}

func NewWorkService(name string) WorkService{
	return &Work{
		WorkName: name,
	}
}


func (w *Work) WorkStart(name string) *Job{
	fmt.Println("start worlk")
	return &Job{
		Name: name,
	}
}
