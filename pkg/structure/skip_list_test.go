package structure

import "testing"


func Test_InsertToSkipList(t *testing.T) {
	list := NewSkipList()
	testCases := []struct {
		Name string
		List *SkipList
		Value []int
	}{
		{
      Name: "Insert at the head",
      Value: []int{5, 6, 7, 10, 2, 4},
    },
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			for _, v := range tc.Value {
				list.Insert(v)
			}
		})
	}
}