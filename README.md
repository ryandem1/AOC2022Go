# aoc-2022-go
Advent of Code 2022 In Go

Detailed problems/prompts can be found here: [Advent of Code 2022](https://adventofcode.com/2022)


## Running a solution
To run any day follow the syntax (replace "3" with whatever day you would like to run):

`make run DAY=3`

This will produce you a result like the following:

```
--------------------------------------------------------day3--------------------------------------------------------

2022/12/04 17:48:45 What is the sum of the priorities of those item types?
2022/12/04 17:48:45 Part 1 Answer:
2022/12/04 17:48:45 >>> 7766 <<<

2022/12/04 17:48:45 What is the sum of the priorities of those item types?
2022/12/04 17:48:45 Part 2 Answer:
2022/12/04 17:48:45 >>> 2415 <<<

--------------------------------------------------------------------------------------------------------------------
```


## Generating a visualization

Sometimes problems can be visualized through monospace textart representations. For those problems, you can run a 
separate "visualize" command to generate the visualization. All visualizations will be saved with the filename:
"visualization.txt".

These often come in handy when in comes to debugging, they can also be kind of cool sometimes.

Small note that if you view the `.txt` files that get outputted with a non-monospace font, the visualizations will not
come out correctly.

To generate a visualization for a day that supports it run:

`make visualize DAY=4`

If it is successful, you should see:

```
go build .
./aoc_2022_go day4 -v > visualization.txt
rm ./aoc_2022_go
```

If the day does not support visualization, you will see:

```
go build .
./aoc_2022_go day3 -v > visualization.txt
panic: No visualization for day3

goroutine 1 [running]:
main.main()
        /Users/ryandemarigny/GolandProjects/AOC2022Go/main.go:69 +0x440
make: *** [visualize] Error 2
```
