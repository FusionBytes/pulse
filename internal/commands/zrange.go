package commands

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

func NewZRange(scoreboards *structure.HashTable) *ZRange {
	return &ZRange{scoreboards: scoreboards}
}

func (z *ZRange) CanDo(method string) bool {
	if method == "zrange" || method == "ZRANGE" {
		return true
	}
	return false
}

func (z *ZRange) Execute(args []string) (interface{}, error) {
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
	scoreboardData, ok := fetchedScoreboard.(*structure.SkipList)
	if !ok {
		return nil, errors.New("can't fetch scoreboard")
	}
	members := scoreboardData.RangeByRank(start, stop)
	var formattedSlice []string

	// Loop through the original slice and format each string
	for i, name := range members {
		formattedSlice = append(formattedSlice, fmt.Sprintf("%d) %s", i+1, name))
	}

	// Join the formatted strings with a newline separator
	result := strings.Join(formattedSlice, "\n")
	return result, nil
}
