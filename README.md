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
"day<dayNumber>-visualization-<visualizationNumber>.txt".

These often come in handy when in comes to debugging, they can also be kind of cool sometimes.

Small note that if you view the `.txt` files that get outputted with a non-monospace font, the visualizations will not
come out correctly.

To generate a visualization for a day that supports it run:

`make visualize DAY=4`

If it is successful, you should see something like:

```
go build .
./aoc_2022_go day4 -visual1 > day4-visualization-1.txt
rm ./aoc_2022_go
```

If the day does not support visualization, you will see something like:

```
go build .
./aoc_2022_go day3 -visual1 > day3-visualization-1.txt
panic: No visualization for day3

goroutine 1 [running]:
main.main()
        /Users/ryandemarigny/GolandProjects/aoc-2022-go/main.go:75 +0x8dc
make: *** [visualize] Error 2
```

### Days with multiple visualizations
Some days might support multiple visualizations. By default, running `make visualize DAY=4` will generate VISUAL #1.

To run a different visualization for a day, run:

`make visualize DAY=5 VISUAL=2`

You can replace the `VISUAL` number with the number of the visualization you want to run. If the visualization number
specified does not exist for the day, you will see the following error:

```
go build .
./aoc_2022_go day5 -visual3 > day5-visualization-3.txt
panic: Only have 2 available visualization(s)

goroutine 1 [running]:
main.main()
        /Users/ryandemarigny/GolandProjects/aoc-2022-go/main.go:127 +0x848
make: *** [visualize] Error 2
```

Note that it will tell you the total amount of visualizations for that day in the error: 
`panic: Only have 2 available visualization(s)`
