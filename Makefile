run:
	go build .
	./aoc_2022_go day${DAY}
	rm ./aoc_2022_go


VISUAL := 1
visualize:
	go build .
	./aoc_2022_go day${DAY} -visual${VISUAL} > day${DAY}-visualization-${VISUAL}.txt
	rm ./aoc_2022_go
