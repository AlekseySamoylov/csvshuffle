package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	println("Type filename and press Enter")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.Replace(fileName, "\n", "", -1)
	wordPairs := ReadCsv(fileName)
	println("Type Q to exit or press ENTER to go to the next word")

	for {
		checkForTheExit(reader)
		russian, english := Shuffle(wordPairs)
		println(russian)
		checkForTheExit(reader)
		println(english)
		println("--------------")
	}

}

func Shuffle(wordPairs map[string]string) (string, string) {

	if wordPairs == nil || len(wordPairs) == 0 {
		return "пустой список", "empty list"
	}
	keys := reflect.ValueOf(wordPairs).MapKeys()
	randomKey := keys[rand.Intn(len(keys))].String()

	return randomKey, wordPairs[randomKey]
}

func ReadCsv(fileName string) map[string]string {
	csvfile, err := os.Open(fileName)
	if err != nil {
		err = nil
		csvfile, err = os.Open(fileName + ".csv")
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}
	}

	r := csv.NewReader(csvfile)
	wordsMap := make(map[string]string)
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		wordsMap[record[0]] = record[1]
	}
	return wordsMap
}

func checkForTheExit(reader *bufio.Reader) {
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	if strings.Compare("q", text) == 0 || strings.Compare("Q", text) == 0 {
		println("See you...")
		os.Exit(0)
	}
}
