package tools

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"unicode"
)

func ThousandSeparator(n int64) string {
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits--
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = '.'
		}
	}
}

func PackageName(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "package ") {
			parts := strings.Fields(line)
			if len(parts) == 2 {
				return parts[1], nil
			}
			return "", fmt.Errorf("invalid package declaration in file: %s", filename)
		}

		return "", fmt.Errorf("no package declaration found in file: %s", filename)
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("empty file: %s", filename)
}

type Counter struct {
	serial int64
	mu     sync.Mutex
}

func (c *Counter) NextSerial() int64 {
	return atomic.AddInt64(&c.serial, 1)
}

func (c *Counter) ResetCounter(value int64) {
	atomic.StoreInt64(&c.serial, value)
}

func DebugMarker(str string) {
	fmt.Println(str + " " + "<<<<<<<<<<<<<< MARK")
}

func ToTitleCase(input string) string {
	words := strings.Fields(input)
	var titledWords []string
	for _, word := range words {
		titleWord := strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		titledWords = append(titledWords, titleWord)
	}
	titleCaseString := strings.Join(titledWords, " ")
	return titleCaseString
}
func SanitizeInput(input string) string {
	dangerousChars := []string{";", "=", "--", "'", "\"", "\\", "/", "*", "%", "select"}
	sanitized := input
	for _, char := range dangerousChars {
		sanitized = strings.ReplaceAll(sanitized, char, "")
	}
	sanitized = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_' {
			return r
		}
		return -1
	}, sanitized)
	return strings.Trim(sanitized, " ")
}
