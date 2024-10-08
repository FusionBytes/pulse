package sortedset

import (
	"errors"
	"fmt"
	"pulse/pkg/structure"
	"strconv"
	"strings"
)

type ZRange struct {
	scoreboards *structure.HashTable
}

func NewZRange() *ZRange {
	return &ZRange{scoreboards: newHashTable()}
}

func (z *ZRange) CanDo(method string) bool {
	if method == "zrange" || method == "ZRANGE" {
		return true
	}
	return false
}

func (z *ZRange) Execute(args []string) (interface{}, error) {
	err := z.validateArgs(args)
	if err != nil {
		return nil, err
	}

	scoreboard := args[0]
	start, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, err
	}
	stop, err := strconv.Atoi(args[2])
	if err != nil {
		return nil, err
	}
	fetchedScoreboard, ok := z.scoreboards.Get(scoreboard)
	if !ok {
		return nil, errors.New("can't find scoreboard")
	}
	scoreboardData, ok := fetchedScoreboard.(*Scoreboard)
	if !ok {
		return nil, errors.New("can't fetch scoreboard")
	}
	members := scoreboardData.skipList.RangeByRank(start, stop)
	var formattedSlice []string

	// Loop through the original slice and format each string
	for i, name := range members {
		formattedSlice = append(formattedSlice, fmt.Sprintf("%d) %s", i+1, name))
	}

	// Join the formatted strings with a newline separator
	result := strings.Join(formattedSlice, "\n")
	return result, nil
}

func (z *ZRange) validateArgs(args []string) error {
	if len(args) < 3 {
		return errors.New("expected at least 2 arguments")
	}
	if args[0] == "" {
		return errors.New("invalid scoreboard")
	}
	if args[1] == "" {
		return errors.New("invalid start rank")
	}
	if args[2] == "" {
		return errors.New("invalid stop rank")
	}
	return nil
}
