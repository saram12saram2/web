package asciiart

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ErrorHandling(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func AsciiArt(word string, banner string) (string, int) {
	StatusCode := 0
	result := ""
	reading_file := banner + ".txt"
	readFile, err := os.Open(reading_file)
	inData, err := ioutil.ReadFile(banner + ".txt")
	if err != nil {
		return "", 404
	}
	if !CheckerHash(inData) {
		return "", 500
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	defer readFile.Close()
	// flag := true
	if len(fileLines) != 855 {
		StatusCode = 500
		return result, StatusCode
	}
	data := make(map[int]string)
	symbol := 32
	for i, line := range fileLines {
		if line == "" {
			data[i+1] = string(rune(symbol))
			symbol++
		}
	}

	lines := []string{"lines", "lines2", "lines3", "lines4", "lines5", "lines6", "lines7", "lines8"}
	// words_replace_all := strings.ReplaceAll(word, "\\n", "\n")
	splitted_words := strings.Split(word, "\n")
	for _, words := range splitted_words {
		for _, letters := range words {
			if letters < 32 || letters > 126 {
				if letters == 13 {
					continue
				}
				StatusCode = 400
				return result, StatusCode
			}
		}
	}
	linesMap := make(map[string]string)
	index_newline := []int{}
	for _, words := range splitted_words {
		count := 0
		for _, lettersInWords := range string(words) {
			for keysIndx, lettersInData := range data {
				if string(lettersInWords) == lettersInData {
					for i := 0; i < len(lines); i++ {
						linesMap[lines[i]] += fileLines[keysIndx+i]
						if i == 0 {
							count += len(fileLines[keysIndx+i])
						}
					}
				}
			}
		}
		index_newline = append(index_newline, count)
	}
	counter_for_zero := 0
	for i := 0; i < len(index_newline); i++ {
		if index_newline[i] == 0 {
			counter_for_zero++
		}
	}
	flag := true
	if counter_for_zero == len(index_newline) {
		index_newline = index_newline[:len(index_newline)-1]
		if counter_for_zero == 3 {
			flag = false
		}
	}
	if counter_for_zero == 3 && !flag {
		result += "\n" + "\n"
	} else {
		start := 0
		if len(index_newline) == 2 && index_newline[0] == 0 && index_newline[1] == 0 {
			result += ""
		} else {
			for z := 0; z < len(index_newline); z++ {
				end := index_newline[z]
				if end == 0 {
					result += "\n"
				} else {
					for i := 0; i < len(lines); i++ {
						if i == len(lines)-1 && z == len(index_newline)-1 {
							result += linesMap[lines[i]][start:start+end] + "\n" + "\n"
						} else {
							result += linesMap[lines[i]][start:start+end] + "\n"
						}
					}
				}
				start = start + end
			}
		}
	}
	return result, StatusCode
}
