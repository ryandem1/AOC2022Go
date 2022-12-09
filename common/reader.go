package common

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

// ReadLinesFromFile will read an input file located in the inputs directory and will send the string lines through a
// channel
func ReadLinesFromFile(fileName string) chan string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(filepath.Join(cwd, "inputs", fileName))

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	lines := make(chan string)
	go func() {
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)

		err = f.Close()
		if err != nil {
			panic(err)
		}
	}()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

// ReadChannelInChunks will break a channel's return into chunks of a certain size. It is best-effort, if there are not
// enough elements to fill up a full chunk, it will return a partial chunk
func ReadChannelInChunks[__T any, __C chan __T](channel __C, chunkSize int) chan []__T {
	chunkedChannel := make(chan []__T)

	go func() {
		var chunk []__T
		for item := range channel {
			chunk = append(chunk, item)
			if len(chunk) == chunkSize {
				chunkedChannel <- chunk

				// Clears chunk and sends to garbage collector. Fun fact, I first tried to just clear the chunk, but
				// keep the memory with chunk = [:0], but that caused a nasty race condition :)
				chunk = nil
			}
		}

		// Sends the last partial chunk if one exists
		if len(chunk) > 0 {
			chunkedChannel <- chunk
		}
		close(chunkedChannel)
	}()
	return chunkedChannel
}

// ReadSliceWithWindow will create a moving window through the slice and yield the window through a channel
func ReadSliceWithWindow[__T any](sl []__T, chunkSize int) chan []__T {
	chunks := make(chan []__T)

	go func() {
		for i := 0; i < len(sl); i++ {
			end := i + chunkSize
			if end > len(sl) {
				end = len(sl)
			}
			chunks <- sl[i:end]
		}
		close(chunks)
	}()
	return chunks
}

// AddChannelBuffer will create a proxy channel for whatever channel is passed and provider a buffer. Good for channels
// that were initially created buffer-less. This is probably a bad design.
func AddChannelBuffer[T any](input chan T, bufferSize int) chan T {
	bufferedOutput := make(chan T, bufferSize)

	go func() {
		for item := range input {
			bufferedOutput <- item
		}
		close(bufferedOutput)
	}()
	return bufferedOutput
}
