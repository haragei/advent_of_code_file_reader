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

func ReadRealFile() []string {
	return readFile("./real.input")
}

func ReadTestFile() []string {
	return readFile("./test.input")
}

func ReadRealFileSingleLine(delimiter string) []string {
	lines := readFile("./real.input")

	return strings.Split(lines[0], delimiter)
}

func ReadTestFileSingleLine(delimiter string) []string {
	lines := readFile("./test.input")

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
