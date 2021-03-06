package stdin

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// StdIn wraps the scanner
type StdIn struct {
	scanner *bufio.Scanner
}

// NewStdIn initialize Stdin using singleton mode
func NewStdIn() *StdIn {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return &StdIn{scanner}
}

// NewStdInLine initialize Stdin using singleton mode
func NewStdInLine() *StdIn {
	scanner := bufio.NewScanner(os.Stdin)
	return &StdIn{scanner}
}

// IsEmpty reports if the In is empty
func (s *StdIn) IsEmpty() bool {
	return !s.scanner.Scan()
}

// ReadString reads the next token or the next line depends on splitfunc, and returns the string.
func (s *StdIn) ReadString() string {
	return s.scanner.Text()
}

// ReadInt reads the next token from this input stream, parses it as a int, and returns the int.
func (s *StdIn) ReadInt() int {
	v, err := strconv.Atoi(s.scanner.Text())
	if err != nil {
		panic(err)
	}
	return v
}

// ReadAllStrings reads all remaining tokens from standard input and returns them as a slice of strings.
func ReadAllStrings() (words []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords) // split by words
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}

// ReadAll reads and returns the remainder of the input, as a string.
func ReadAll() string {
	var s strings.Builder
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		s.WriteString(line)
	}
	return s.String()
}
