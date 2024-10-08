package sortedset

import (
	"errors"
	"pulse/pkg/structure"
)

type ZScore struct {
	scoreboards *structure.HashTable
}

func NewZScore() *ZScore {
	return &ZScore{scoreboards: newHashTable()}
}

func (z *ZScore) CanDo(method string) bool {
	if method == "zscore" || method == "ZSCORE" {
		return true
	}
	return false
}

func (z *ZScore) Execute(args []string) (interface{}, error) {
	err := z.validateArgs(args)
	if err != nil {
		return nil, err
	}

	scoreboard := args[0]
	member := args[1]
	fetchedScoreboard, ok := z.scoreboards.Get(scoreboard)
	if !ok {
		return nil, errors.New("can't find scoreboard")
	}
	scoreboardData, ok := fetchedScoreboard.(*Scoreboard)
	if !ok {
		return nil, errors.New("can't fetch scoreboard")
	}
	score, ok := scoreboardData.hashMap.Get(member)
	if !ok {
		return nil, errors.New("can't find member")
	}
	return score, nil
}

func (z *ZScore) validateArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("expected at least 2 arguments")
	}
	if args[0] == "" {
		return errors.New("invalid scoreboard")
	}
	if args[1] == "" {
		return errors.New("invalid member")
	}
	return nil
}
