package reader

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ReadFile(filename string) []string {
	return readFile(filename)
}

func ReadFileRunes(filename string) [][]rune {
	return readFileRunes(filename)
}

func ReadFileSingleLine(filename string, delimiter string) []string {
	lines := readFile(filename)

	return strings.Split(lines[0], delimiter)
}

func readFile(fileName string) (lines []string) {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	lineScanner := bufio.NewScanner(file)
	for lineScanner.Scan() {
		lines = append(lines, lineScanner.Text())
	}

	check(lineScanner.Err())
	return
}

func readFileRunes(fileName string) (lines [][]rune) {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	lineScanner := bufio.NewScanner(file)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		lines = append(lines, []rune(line))
	}

	check(lineScanner.Err())
	return
}
