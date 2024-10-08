package sortedset

import (
	"errors"
	"pulse/pkg/structure"
)

type ZRank struct {
	scoreboards *structure.HashTable
}

func NewZRank() *ZRank {
	return &ZRank{scoreboards: newHashTable()}
}

func (z *ZRank) CanDo(method string) bool {
	if method == "zrank" || method == "ZRANK" {
		return true
	}
	return false
}

func (z *ZRank) Execute(args []string) (interface{}, error) {
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
	rank, err := scoreboardData.skipList.Rank(member)
	if err != nil {
		return nil, nil
	}

	return rank, err
}

func (z *ZRank) validateArgs(args []string) error {
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
