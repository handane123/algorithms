package str

import (
	"fmt"

	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
	"github.com/handane123/algorithms/digraph"
)

// NFA struct provides a data type for creating a nondeterministic finite state automaton (NFA) from
// a regular expression and testing whether a given string is matched by that regular expression.
// It supports the following operations: concatenation, closure, binary or, and parentheses.
type NFA struct {
	graph  *digraph.Digraph // digraph of  ε-transitions
	regexp string           // regular expression
	m      int              // number of characters in regular expression
}

// NewNFA initializes the NFA from the specified regular expression.
func NewNFA(regexp string) *NFA {
	nfa := &NFA{regexp: regexp, m: len(regexp), graph: digraph.NewDigraph(len(regexp) + 1)}
	ops := arraystack.New()
	for i := 0; i < nfa.m; i++ {
		lp := i
		if string(regexp[i]) == "(" || string(regexp[i]) == "|" {
			ops.Push(i)
		} else if string(regexp[i]) == ")" {
			val, _ := ops.Pop()
			or := val.(int)

			if string(regexp[or]) == "|" {
				val, _ := ops.Pop()
				lp = val.(int)
				nfa.graph.AddEdge(lp, or+1)
				nfa.graph.AddEdge(or, i)
			} else if string(regexp[or]) == "(" {
				lp = or
			}
		}

		if i < nfa.m-1 && string(regexp[i+1]) == "*" {
			nfa.graph.AddEdge(lp, i+1)
			nfa.graph.AddEdge(i+1, lp)
		}
		if string(regexp[i]) == "(" || string(regexp[i]) == "*" || string(regexp[i]) == ")" {
			nfa.graph.AddEdge(i, i+1)
		}
	}
	if ops.Size() != 0 {
		panic("invalid regular expression")
	}
	return nfa
}

// Recognizes returns true if the text is matched by the regular expression.
func (nfa *NFA) Recognizes(txt string) bool {
	dfs := digraph.NewDirectedDFS(nfa.graph, 0)
	var pc []int
	//states reachable from start by ε-transitions
	for v := 0; v < nfa.graph.V(); v++ {
		if dfs.IsMarked(v) {
			pc = append(pc, v)
		}
	}

	// Compute possible NFA states for txt[i+1]
	for i := range txt {
		if string(txt[i]) == "*" || string(txt[i]) == "|" || string(txt[i]) == "(" || string(txt[i]) == ")" {
			panic(fmt.Sprintf("text contains the metacharacter '%c'\n", txt[i]))
		}
		states := make([]int, 0) // set of states reachable after scanning past txt[i]
		for _, v := range pc {
			if v == nfa.m {
				continue // not necessarily a match (RE needs to match full text)
			}
			if nfa.regexp[v] == txt[i] || string(nfa.regexp[v]) == "." {
				states = append(states, v+1)
			}
		}
		// follow ε-transitions
		dfs = digraph.NewDirectedDFSources(nfa.graph, states)
		pc = make([]int, 0)
		for v := 0; v < nfa.graph.V(); v++ {
			if dfs.IsMarked(v) {
				pc = append(pc, v)
			}
		}

		if len(pc) == 0 { // no states reachable
			return false
		}
	}

	// check for accept state
	for _, v := range pc {
		if v == nfa.m {
			return true // accept if can end in state M
		}
	}
	return false
}