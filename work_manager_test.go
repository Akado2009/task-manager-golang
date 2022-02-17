package main

import (
	"math/rand"
	"testing"
)

var trs int = 10

const max = 100
const min = 1

func calcFNTesting(element WET) WET {
	return 2 * element
}

func getTestData(size int) *[]WET {
	data := make([]WET, size)
	for i := 0; i < size; i++ {
		data[i] = WET(rand.Intn(max - min))
	}
	return &data
}

func TestNewWorkManager(t *testing.T) {
	s := &WorkManager{
		threshold: trs,
	}
	sFN := NewWorkManager(trs)
	if s.threshold != sFN.threshold {
		t.Errorf("Got different thresholds. Direct: %d, using function: %d", s.threshold, sFN.threshold)
	}
}

func TestSingleThread(t *testing.T) {
	data := getTestData(trs - 1)

	wm := NewWorkManager(trs)
	res := wm.RunJob(data, calcFNTesting, false)

	if len(*data) != len(*res) {
		t.Errorf("Got different lengths. Input: %d. Output: %d", len(*data), len(*res))
	}
	for i, el := range *res {
		actual := calcFNTesting((*data)[i])
		if el != actual {
			t.Errorf("Got different results at index %d. Input: %d. Output: %d", i, el, actual)
		}
	}
}

func TestMultipleThread(t *testing.T) {
	data := getTestData(2 * trs)

	wm := NewWorkManager(trs)
	res := wm.RunJob(data, calcFNTesting, false)

	if len(*data) != len(*res) {
		t.Errorf("Got different lengths. Input: %d. Output: %d", len(*data), len(*res))
	}
	for i, el := range *res {
		actual := calcFNTesting((*data)[i])
		if el != actual {
			t.Errorf("Got different results at index %d. Input: %d. Output: %d", i, el, actual)
		}
	}
}

func TestAllThreads(t *testing.T) {
	data := getTestData(2 * trs)

	wm := NewWorkManager(trs)
	res := wm.RunJob(data, calcFNTesting, true)

	if len(*data) != len(*res) {
		t.Errorf("Got different lengths. Input: %d. Output: %d", len(*data), len(*res))
	}
	for i, el := range *res {
		actual := calcFNTesting((*data)[i])
		if el != actual {
			t.Errorf("Got different results at index %d. Input: %d. Output: %d", i, el, actual)
		}
	}
}
