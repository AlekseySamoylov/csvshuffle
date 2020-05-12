package main

import "testing"

var testCases = []struct {
	lines []Line
}{
	{
		[]Line{Line{[]string{"с утра до вечера", "all day long"}}, Line{[]string{"так же, как и …", "as well as"}}},
	},
	{
		[]Line{},
	},
	{
		nil,
	},
}

func TestShuffle(t *testing.T) {
	// Given
	for _, testCase := range testCases {
		// WhenR
		line := Shuffle(testCase.lines)

		if len(line.elements) == 0 || line.elements[0] == "" {
			t.Errorf("Test failed input: A = %s , results = %s == %s", testCase.lines, line.elements[0], line.elements[1])
		}
	}
}

var csvTestCases = []struct {
	fileName string
}{
	{
		"test.csv",
	},
	{
		"test",
	},
}

func TestReadCsv(t *testing.T) {

	for _, testCase := range csvTestCases {

		csvFile := testCase.fileName
		expectedKey := "первый"
		wordPairs := ReadCsv(csvFile)
		if len(wordPairs) != 2 {
			t.Errorf("Test failed: A = %s , map = %s doesn't contain %s", csvFile, wordPairs, expectedKey)
		}

	}
}
