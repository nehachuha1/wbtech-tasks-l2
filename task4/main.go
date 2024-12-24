package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type FileFlags struct {
	Filename  string
	Column    int
	HasNumber bool
	IsReverse bool
	IsUnique  bool
}

func ParseFlags() *FileFlags {
	f := &FileFlags{}

	flag.StringVar(&f.Filename, "f", "", "Filename")
	flag.IntVar(&f.Column, "col", -1, "Column that should be sorted")
	flag.BoolVar(&f.HasNumber, "n", false, "Sort a number value")
	flag.BoolVar(&f.IsReverse, "r", false, "Reverse sort")
	flag.BoolVar(&f.IsUnique, "u", false, "Sort only unique strings")

	return f
}

func MakeSortByColumns(lines []string, col int) []string {
	m := make(map[string]string)

	for i := 0; i < len(lines); i++ {
		splittedLine := strings.Split(lines[i], " ")
		m[splittedLine[col]] = lines[i]
	}

	keys := GetKeys(m)
	sort.Strings(keys)

	sortedLines := make([]string, len(lines))
	for i, key := range keys {
		sortedLines[i] = m[key]
	}

	return sortedLines
}

func MakeSort(lines []string, f *FileFlags) []string {
	if f.Column >= 0 {
		MakeSortByColumns(lines, f.Column)
	} else {
		sort.Strings(lines)
	}

	if f.IsReverse {
		ReverseData(lines)
	}

	if f.IsUnique {
		lines = Unique(lines)
	}

	return lines
}

func GetKeys(m map[string]string) []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func ReverseData(data []string) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func Unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func ReadFile(path string) []string {
	fileRows := make([]string, 0)

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error by opening file: %v", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatalf("Error by closing file: %v", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileRows = append(fileRows, scanner.Text())
	}

	return fileRows
}

func WriteFile(path string, data []string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("Error by creating new file: %v", err)
	}

	for i := 0; i < len(data); i++ {
		_, err = fmt.Fprintf(file, data[i])
		if err != nil {
			log.Fatalf("Error by writing to the file: %v", err)
		}
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatalf("Failed to close file: %v", err)
		}
	}()
}

func main() {
	outputFile := "Projects/wbtech-tasks-l2/task4"

	flags := ParseFlags()
	lines := ReadFile(flags.Filename)
	sortedLines := MakeSort(lines, flags)
	WriteFile(outputFile, sortedLines)
}
