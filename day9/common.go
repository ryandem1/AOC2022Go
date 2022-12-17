package day9

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"strconv"
	"strings"
)

func readMotions() chan *ropeMotion {
	motions := make(chan *ropeMotion)

	go func() {
		for line := range common.ReadLinesFromFile("day9") {
			fields := strings.Fields(line)

			direction := map[string]motionDirection{
				"U": up,
				"D": down,
				"L": left,
				"R": right,
			}[fields[0]]
			amount, err := strconv.Atoi(fields[1])
			if err != nil {
				panic(err)
			}

			motion := &ropeMotion{
				direction: direction,
				amount:    amount,
			}
			motions <- motion
		}
		close(motions)
	}()
	return motions
}
