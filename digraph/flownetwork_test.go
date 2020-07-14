package digraph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/stretchr/testify/assert"
)

func TestFlowNetWork(t *testing.T) {
	assert := assert.New(t)

	assert.PanicsWithValue("number of vertices in a graph must be non-negative", func() { NewFlowNetwork(-1) })
	fn := NewFlowNetwork(2)
	assert.Equal(2, fn.V())
	assert.Equal(0, fn.E())

	assert.PanicsWithValue("number of edges must be non-negative", func() { NewFlowNetworkVE(3, -1) })
	fn1 := NewFlowNetworkVE(3, 1)
	assert.Equal(3, fn1.V())

	tinyfn := "6\n" +
		"8\n" +
		"0 1 2.0\n" +
		"0 2 3.0\n" +
		"1 3 3.0\n" +
		"1 4 1.0\n" +
		"2 3 1.0\n" +
		"2 4 1.0\n" +
		"3 5 2.0\n" +
		"4 5 3.0\n"
	buf := strings.NewReader(tinyfn)
	s := bufio.NewScanner(buf)
	s.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: s}

	fn2 := NewFlowNetworkIn(in)
	assert.Equal(6, fn2.V())
	assert.Equal(8, fn2.E())
	adjedges := []*FlowEdge{
		NewFlowEdge(0, 2, 3.0),
		NewFlowEdge(0, 1, 2.0),
	}
	assert.Equal(adjedges, fn2.Adj(0))

	edges := []*FlowEdge{
		NewFlowEdge(0, 2, 3.0),
		NewFlowEdge(0, 1, 2.0),
		NewFlowEdge(1, 4, 1.0),
		NewFlowEdge(1, 3, 3.0),
		NewFlowEdge(2, 4, 1.0),
		NewFlowEdge(2, 3, 1.0),
		NewFlowEdge(3, 5, 2.0),
		NewFlowEdge(4, 5, 3.0),
	}
	assert.Equal(edges, fn2.Edges())

	toString := "6 8\n" +
		"0: 0->2 0.000000/3.000000  0->1 0.000000/2.000000  \n" +
		"1: 1->4 0.000000/1.000000  1->3 0.000000/3.000000  \n" +
		"2: 2->4 0.000000/1.000000  2->3 0.000000/1.000000  \n" +
		"3: 3->5 0.000000/2.000000  \n" +
		"4: 4->5 0.000000/3.000000  \n" +
		"5: \n"

	assert.Equal(toString, fn2.String())
	assert.Panics(func() { fn2.Adj(10) })

	tinyfn1 := "6\n" +
		"-1\n" +
		"0 1 2.0\n" +
		"0 2 3.0\n" +
		"1 3 3.0\n" +
		"1 4 1.0\n" +
		"2 3 1.0\n" +
		"2 4 1.0\n" +
		"3 5 2.0\n" +
		"4 5 3.0\n"
	buf1 := strings.NewReader(tinyfn1)
	s1 := bufio.NewScanner(buf1)
	s1.Split(bufio.ScanWords)
	in1 := &stdin.In{Scanner: s1}

	assert.PanicsWithValue("number of edges must be non-negative", func() { NewFlowNetworkIn(in1) })

}
