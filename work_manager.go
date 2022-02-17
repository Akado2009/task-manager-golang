package main

import (
	"sync"
)

// WET - WorkElement Type
type WET int
type WETFunction func(WET) WET

type WorkManager struct {
	threshold int
}

// NewWorkManager creates a new WorkManager instance
func NewWorkManager(threshold int) *WorkManager {
	return &WorkManager{
		threshold: threshold,
	}
}

// RunJob is a general purpose function to run
// an average WETFunction with []WET as input
// Result is returned from this function
func (wm *WorkManager) RunJob(input *[]WET, fn WETFunction, runSeparately bool) *[]WET {
	// check whether the input <> threshold
	var result *[]WET
	if len(*input) > wm.threshold {
		if runSeparately {
			result = wm.dispatchAllThreads(input, fn)
		} else {
			result = wm.dispatchMultipleThread(input, fn)
		}
	} else {
		result = wm.dispatchSingleThread(input, fn)
	}

	return result
}

func (wm *WorkManager) dispatchSingleThread(input *[]WET, fn WETFunction) *[]WET {
	r := make([]WET, len(*input))
	for i, element := range *input {
		r[i] = fn(element)
	}
	return &r
}

func (wm *WorkManager) dispatchMultipleThread(input *[]WET, fn WETFunction) *[]WET {
	r := make([]WET, len(*input))
	var wg sync.WaitGroup
	chunks := wm.splitData(input)
	for i, chunk := range chunks {
		wg.Add(1)
		go func(ch *[]WET, index int) {
			defer wg.Done()
			for j, element := range *ch {
				r[index*wm.threshold+j] = fn(element)
			}
		}(chunk, i)
	}
	wg.Wait()
	return &r
}

func (wm *WorkManager) dispatchAllThreads(input *[]WET, fn WETFunction) *[]WET {
	r := make([]WET, len(*input))
	var wg sync.WaitGroup
	for i, element := range *input {
		wg.Add(1)
		go func(el WET, index int) {
			defer wg.Done()
			r[index] = fn(el)
		}(element, i)
	}
	wg.Wait()
	return &r
}

func (wm *WorkManager) splitData(input *[]WET) []*[]WET {
	chunks := make([]*[]WET, 0)
	chunkSize := wm.threshold
	for i := 0; i < len(*input); i += chunkSize {
		end := i + chunkSize
		if end > len(*input) {
			end = len(*input)
		}
		ch := (*input)[i:end]
		chunks = append(chunks, &ch)
	}

	return chunks
}
