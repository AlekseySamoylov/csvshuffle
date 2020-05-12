package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
)

// var delim byte = '\r'
var delim byte = '\n'
var delimString = "\n"

func main() {
	reader := bufio.NewReader(os.Stdin)
	println("Type filename and press Enter")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.Replace(fileName, delimString, "", -1)
	wordPairs := ReadCsv(fileName)
	println("Type Q to exit or S to shuffle or press ENTER to go to the next word")

	for {
		for idx := 0; idx < len(wordPairs); idx++ {
			line := wordPairs[idx]
			words := line.elements
			for subIdx := 0; subIdx < len(words); subIdx++ {
				println(words[subIdx])
				if checkForTheExit(reader) {
					println("SHUFFLE")
					rand.Shuffle(len(wordPairs), func(i, j int) { wordPairs[i], wordPairs[j] = wordPairs[j], wordPairs[i] })
					println(wordPairs[0].elements[0])
					idx = 0
					subIdx = 0
					continue
				}
			}
			println("--------------")
		}
		println("NEXT ROUND")
	}

}

type Line struct {
	elements []string
}

func Shuffle(lines []Line) Line {

	if lines == nil || len(lines) == 0 {
		return Line{[]string{"пустой список | empty list"}}
	}

	return lines[rand.Intn(len(lines))]
}

func ReadCsv(fileName string) []Line {
	csvfile, err := os.Open(fileName)
	if err != nil {
		err = nil
		csvfile, err = os.Open(fileName + ".csv")
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}
	}

	r := csv.NewReader(csvfile)
	var words []Line
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		words = append(words, Line{record})
	}
	return words
}

func checkForTheExit(reader *bufio.Reader) bool {
	text, _ := reader.ReadString(delim)
	// convert CRLF to LF
	text = strings.Replace(text, delimString, "", -1)
	if strings.Compare("q", text) == 0 || strings.Compare("Q", text) == 0 {
		println("See you...")
		os.Exit(0)
	}
	return strings.Compare("s", text) == 0 || strings.Compare("S", text) == 0
}
