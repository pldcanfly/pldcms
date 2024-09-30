package flow

import (
	"context"
	"fmt"
	"net/http"
)

type Flow struct {
	Ctx      context.Context
	Request  *http.Request
	Response *http.Response
	View     string
	Flowers  []Flower
}

type Flower interface {
	Run() (bool, error)
	Init()
	GetName() string
}

func NewFlow() *Flow {
	return &Flow{
		Ctx: context.Background(),
	}
}

func (f *Flow) AddFlower(fl Flower) {
	f.Flowers = append(f.Flowers, fl)
	fmt.Println("Added Flower: ", fl.GetName())
}

func (f *Flow) Run(req *http.Request, res *http.Response) {
	f.Request = req
	f.Response = res
	for _, fl := range f.Flowers {
		next, err := fl.Run()
		if err != nil {
			fmt.Errorf("%w", err)
		}
		if !next {
			break
		}
	}
}
