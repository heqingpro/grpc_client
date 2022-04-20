package main

import (
	"fmt"
	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
	"helloword/mock"
	"testing"
)

var jobService *JobService

func init() {
	fmt.Println("test init")
	ser := mock.NewWorkService("unit_test")
	jobService = NewJobService(ser)
}

func TestRun(t *testing.T) {
	Convey("test Run", t, func() {
		Convey("name1", func() {
			So(jobService.Run("name1"), ShouldEqual, "name1") 
		})
		Convey("name2", func() {
			jobService.Run("name2")
			So(jobService.Run("name2"), ShouldEqual, "name2") 
		})
	})
}

func TestRun2(t *testing.T) {
	Convey("test Run2", t, func(){
		// Convey("test1", func()  {
		// 	jobService = new(JobService)
		// 	jobService.workService = mock.NewWorkService("")
		// 	patches := ApplyPrivateMethod(jobService, "workService.WorkStart", func (JobService) *mock.Job {
		// 		return &mock.Job{
		// 			Name: "workStaar",
		// 		}
		// 	})
		// 	defer patches.Reset()
		// 	So(jobService.Run("testRun2"), ShouldEqual, "WorkStart")
		// })
		Convey("test1", func()  {
			jobService = new(JobService)
			workService := mock.NewWorkService("")
			patches := ApplyMethodReturn(workService, "WorkStart", &mock.Job{
				Name: "workStaar",
			})
			jobService.workService = workService
			defer patches.Reset()
			So(jobService.Run("testRun2"), ShouldEqual, "WorkStart")
		})
	})
}