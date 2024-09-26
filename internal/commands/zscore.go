package commands

import (
	"errors"
	"pulse/pkg/structure"
)

type ZScore struct {
	scoreboards *structure.HashTable
}

func NewZScore(scoreboards *structure.HashTable) *ZScore {
	return &ZScore{scoreboards: scoreboards}
}

func (z *ZScore) CanDo(method string) bool {
	if method == "zscore" || method == "ZSCORE" {
		return true
	}
	return false
}

func (z *ZScore) Execute(args []string) (interface{}, error) {
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
	score, ok := scoreboardData.Score(member)
	if !ok {
		return nil, errors.New("can't find member")
	}
	return score, nil
}
