package day4

const TotalSections = 100 // We will only see sections with IDs < 100. Really only import for the visualization

// a cleaningAssignment is a range of camp section IDs that a single Elf is assigned to clean
type cleaningAssignment struct {
	start int // Starting section ID
	end   int // Ending section ID
}
