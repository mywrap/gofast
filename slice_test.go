package gofast

import (
	"encoding/json"
	"testing"
)

func TestIndexInts(t *testing.T) {
	if IndexInts([]int{1, 6, 8, 9, 3, 7}, 8) != 2 {
		t.Error()
	}
	if IndexInts([]int{1, 6, 8, 9, 3, 7}, 10) != -1 {
		t.Error()
	}
}

func TestIndexStrings(t *testing.T) {
	if IndexStrings([]string{"1", "6", "8", "9", "3", "7"}, "8") != 2 {
		t.Error()
	}
	if IndexStrings([]string{"1", "6", "8", "9", "3", "7"}, "10") != -1 {
		t.Error()
	}
}

func TestSplitToBatches(t *testing.T) {
	for _, test := range []struct {
		slice     []string
		batchSize int
		batches   [][]string
	}{
		{slice: nil, batchSize: 3,
			batches: [][]string{}},
		{slice: []string{"0", "1", "2"}, batchSize: 3,
			batches: [][]string{{"0", "1", "2"}}},
		{slice: []string{"0", "1", "2", "3"}, batchSize: 3,
			batches: [][]string{{"0", "1", "2"}, {"3"}}},
		{slice: []string{"0", "1", "2", "3", "4", "5"}, batchSize: 3,
			batches: [][]string{{"0", "1", "2"}, {"3", "4", "5"}}},
	} {
		realBatches := SplitToBatches(test.slice, test.batchSize)
		//t.Logf("realBatches: %#v", realBatches)
		//t.Logf("expeBatches: %#v", test.batches)
		j1, _ := json.Marshal(realBatches)
		j2, _ := json.Marshal(test.batches)
		if string(j1) != string(j2) {
			t.Errorf("error: real: %s, expected: %s", j1, j2)
		}
	}
}
