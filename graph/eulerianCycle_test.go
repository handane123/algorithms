package graph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestEulerianCycle(t *testing.T) {
	assert := assert.New(t)

	graphdata1 := "3\n" +
		"4\n" +
		"0 1\n" +
		"1 2\n" +
		"0 0\n" +
		"2 0"
	buf1 := strings.NewReader(graphdata1)
	scanner1 := bufio.NewScanner(buf1)
	scanner1.Split(bufio.ScanWords)
	in1 := &stdin.In{Scanner: scanner1}
	g1, err1 := NewGraphIn(in1)
	assert.Nil(err1)
	ec1 := NewEulerianCycle(g1)
	assert.True(ec1.HasEulerianCycle())
	assert.Equal([]int{0, 2, 1, 0, 0}, ec1.GetCycle())

	// graph with zero edge
	G2 := NewGraph(3)
	ec2 := NewEulerianCycle(G2)
	assert.Nil(ec2)

	// not all vertices have even degree
	graphdata3 := "3\n" +
		"2\n" +
		"0 1\n" +
		"0 2\n"
	buf3 := strings.NewReader(graphdata3)
	scanner3 := bufio.NewScanner(buf3)
	scanner3.Split(bufio.ScanWords)
	in3 := &stdin.In{Scanner: scanner3}
	g3, err3 := NewGraphIn(in3)
	assert.Nil(err3)
	ec3 := NewEulerianCycle(g3)
	assert.Nil(ec3)

	// graphdata4 := "1\n" +
	// 	"1\n" +
	// 	"0 0\n"
	// buf4 := strings.NewReader(graphdata4)
	// scanner4 := bufio.NewScanner(buf4)
	// scanner4.Split(bufio.ScanWords)
	// in4 := &stdin.In{Scanner: scanner4}
	// g4, err4 := NewGraphIn(in4)
	// assert.Nil(err4)
	// ec4 := NewEulerianCycle(g4)
	// assert.False(ec4.HasEulerianCycle())

}