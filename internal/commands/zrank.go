package commands

import (
	"errors"
	"pulse/pkg/structure"
)

type ZRank struct {
	scoreboards *structure.HashTable
}

func NewZRank(scoreboards *structure.HashTable) *ZRank {
	return &ZRank{scoreboards: scoreboards}
}

func (z *ZRank) CanDo(method string) bool {
	if method == "zrank" || method == "ZRANK" {
		return true
	}
	return false
}

func (z *ZRank) Execute(args []string) (interface{}, error) {
	scoreboard := args[0]
	member := args[1]
	fetchedScoreboard, ok := z.scoreboards.Get(scoreboard)
	if !ok {
		return nil, errors.New("can't find scoreboard")
	}
	scoreboardData, ok := fetchedScoreboard.(*structure.SkipList)
	if !ok {
		return nil, errors.New("can't fetch scoreboard")
	}
	rank := scoreboardData.Rank(member)
	return rank, nil
}
