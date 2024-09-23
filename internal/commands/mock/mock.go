package mock

import "fmt"


type MockCommand struct {
	
}

func (m *MockCommand) CanDo(method string) bool {
	fmt.Println(method)
	return true
}

func (m *MockCommand) Execute(args []string) (interface{}, error) {
	fmt.Printf("the args is %+v \n", args)
	return "success", nil
}