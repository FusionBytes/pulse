package sortedset

import (
	"errors"
	"pulse/pkg/structure"
	"strconv"
)

type ZAdd struct {
	scoreboards *structure.HashTable
}

func NewZAdd() *ZAdd {
	return &ZAdd{scoreboards: newHashTable()}
}

func (z *ZAdd) CanDo(method string) bool {
	if method == "zadd" || method == "ZADD" {
		return true
	}
	return false
}

func (z *ZAdd) Execute(args []string) (interface{}, error) {
	err := z.validateArgs(args)
	if err != nil {

	}
	scoreboard := args[0]
	score, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, err
	}
	member := args[2]
	fetchedScoreboard, ok := z.scoreboards.Get(scoreboard)
	if !ok {
		fetchedScoreboard = &Scoreboard{
			skipList: structure.NewSkipList(),
			hashMap:  structure.NewHashTable(8, 0.75),
		}
		err := z.scoreboards.Insert(scoreboard, fetchedScoreboard)
		if err != nil {
			return nil, err
		}
	}
	scoreboardData, ok := fetchedScoreboard.(*Scoreboard)
	if !ok {
		return nil, errors.New("can't fetch scoreboard")
	}

	err = scoreboardData.hashMap.Insert(member, score)
	if err != nil {
		return nil, err
	}

	scoreboardData.skipList.Insert(score, member)

	return "success", nil
}

func (z *ZAdd) validateArgs(args []string) error {
	if len(args) < 3 {
		return errors.New("expected at least 2 arguments")
	}
	if args[0] == "" {
		return errors.New("invalid scoreboard")
	}
	if args[1] == "" {
		return errors.New("invalid score")
	}
	if args[2] == "" {
		return errors.New("invalid member")
	}
	return nil
}
