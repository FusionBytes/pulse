package commands

import (
	"errors"
	"pulse/pkg/structure"
	"strconv"
)

type ZAdd struct {
	scoreboards *structure.HashTable
}

func NewZAdd(scoreboards *structure.HashTable) *ZAdd {
	return &ZAdd{scoreboards: scoreboards}
}

func (z *ZAdd) CanDo(method string) bool {
	if method == "zadd" || method == "ZADD" {
		return true
	}
	return false
}

func (z *ZAdd) Execute(args []string) (interface{}, error) {
	scoreboard := args[0]
	score, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, err
	}
	member := args[2]
	fetchedScoreboard, ok := z.scoreboards.Get(scoreboard)
	if !ok {
		fetchedScoreboard = structure.NewSkipList()
		err := z.scoreboards.Insert(scoreboard, fetchedScoreboard)
		if err != nil {
			return nil, err
		}
	}
	scoreboardData, ok := fetchedScoreboard.(*structure.SkipList)
	if !ok {
		return nil, errors.New("can't fetch scoreboard")
	}

	scoreboardData.Insert(score, member)
	return "success", nil
}
