package structure

import "testing"

func Test_InsertToSkipList(t *testing.T) {
	list := NewSkipList()
	testCases := []struct {
		Name  string
		List  *SkipList
		Value []struct {
			score  int
			member string
		}
	}{
		{
			Name: "Insert at the head",
			Value: []struct {
				score  int
				member string
			}{
				{1, "member1"},
				{2, "member2"},
				{3, "member3"},
				{4, "member4"},
				{5, "member5"},
				{6, "member6"},
				{7, "member7"},
				{8, "member8"},
				{9, "member9"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			for _, v := range tc.Value {
				list.Insert(v.score, v.member)
			}
		})
	}
}
