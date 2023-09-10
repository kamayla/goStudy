package main

import "fmt"

type OrderService interface {
	Apply(int) error
}

type ServiceImpl struct{}

func (c *ServiceImpl) Apply(id int) error {
	fmt.Println(id)
	return nil
}

type Application struct {
	OrderService
}

func (app *Application) Run(id int) error {
	return app.Apply(id)
}

func main() {
	app := Application{
		OrderService: &ServiceImpl{},
	}

	_ = app.Run(19)
}
