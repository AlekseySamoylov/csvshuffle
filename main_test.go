package main

import "testing"

var testCases = []struct {
	wordPairs map[string]string
}{
	{
		map[string]string{"с утра до вечера": "all day long", "так же, как и …": "as well as"},
	},
	{
		map[string]string{},
	},
	{
		nil,
	},
}

func TestShuffle(t *testing.T) {
	// Given
	for _, testCase := range testCases {
		// WhenR
		russian, english := Shuffle(testCase.wordPairs)

		if (testCase.wordPairs[russian] != english && russian != "пустой список") || russian == "" {
			t.Errorf("Test failed input: A = %s , results = %s == %s", testCase.wordPairs, english, russian)
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
