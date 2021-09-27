package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"goMicroCode/message"
	"time"
)

func main() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Name("student.client"),
	)

	service.Init(micro.Registry(reg))

	studentService := message.NewStudentServiceClient("student_service", service.Client())

	res, err := studentService.GetStudent(context.TODO(), &message.StudentRequest{Name: "jack"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Name)
	fmt.Println(res.Classes)
	fmt.Println(res.Grade)
	time.Sleep(50 * time.Second)
}
