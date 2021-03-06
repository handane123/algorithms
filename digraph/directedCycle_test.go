package digraph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/stretchr/testify/assert"
)

func TestDirectedCycle(t *testing.T) {
	assert := assert.New(t)
	tinyDG := "13\n" +
		"22\n" +
		"4  2\n" +
		"2  3\n" +
		"3  2\n" +
		"6  0\n" +
		"0  1\n" +
		"2  0\n" +
		"11 12\n" +
		"12  9\n" +
		"9 10\n" +
		"9 11\n" +
		"7  9\n" +
		"10 12\n" +
		"11  4\n" +
		"4  3\n" +
		"3  5\n" +
		"6  8\n" +
		"8  6\n" +
		"5  4\n" +
		"0  5\n" +
		"6  4\n" +
		"6  9\n" +
		"7  6\n"

	buf := strings.NewReader(tinyDG)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: scanner}
	g, err := NewDigraphIn(in)
	assert.Nil(err)

	finder := NewDirectedCycle(g)
	assert.True(finder.HasCycle())
	assert.Equal([]int{3, 5, 4, 3}, finder.GetCycle())

	tinyDG1 := "2\n" +
		"1\n" +
		"0 1\n"

	buf1 := strings.NewReader(tinyDG1)
	scanner1 := bufio.NewScanner(buf1)
	scanner1.Split(bufio.ScanWords)
	in1 := &stdin.In{Scanner: scanner1}
	g1, err1 := NewDigraphIn(in1)
	assert.Nil(err1)

	finder1 := NewDirectedCycle(g1)
	assert.False(finder1.HasCycle())
	assert.Equal([]int(nil), finder1.GetCycle())
}
