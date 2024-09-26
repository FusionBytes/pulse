package mock

import "fmt"

type Mock struct {
}

func NewMock() *Mock {
	return &Mock{}
}

func (m *Mock) CanDo(method string) bool {
	fmt.Println(method)
	return true
}

func (m *Mock) Execute(args []string) (interface{}, error) {
	fmt.Printf("the args is %+v \n", args)
	return "success", nil
}
