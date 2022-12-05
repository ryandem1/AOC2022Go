run:
	go build .
	./aoc_2022_go day${DAY}
	rm ./aoc_2022_go

visualize:
	go build .
	./aoc_2022_go day${DAY} -v > visualization.txt
	rm ./aoc_2022_go
