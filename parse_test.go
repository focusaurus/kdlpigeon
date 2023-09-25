package kdl

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mustRootNode(result interface{}) Node {
	if s1, ok := result.([]interface{}); ok {
		if s2, ok := s1[1].([]interface{}); ok {
			return s2[0].(Node)
		}
	}
	panic("Root node not found")
}

func emptyNode() Node {
	return Node{
		Children: []Node{},
		Props:    []Prop{},
		Values:   []Value{},
	}
}

func TestParse(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected func() Node
	}{
		"just identifier": {
			input: "foo",
			expected: func() Node {
				expected := emptyNode()
				expected.Identifier = "foo"
				return expected
			},
		},
		"simple values": {
			input: `foo 0 3.14 "hi" null`,
			expected: func() Node {
				expected := emptyNode()
				expected.Identifier = "foo"
				expected.Values = []Value{
					parseValue(0.0),
					parseValue(3.14),
					parseValue("hi"),
					parseValue(nil),
				}
				return expected
			},
		},
		"!simple props": {
			input: `foo a=0 b=3.14 c="hi" d=null`,
			expected: func() Node {
				expected := emptyNode()
				expected.Identifier = "foo"
				expected.Props = []Prop{
					{Identifier: "a", Value: parseValue(0.0)},
					{Identifier: "b", Value: parseValue(3.14)},
					{Identifier: "c", Value: parseValue("hi")},
					{Identifier: "d", Value: parseValue(nil)},
				}
				return expected
			},
		},
	}

	// To debug a particular test case first, add ! as a prefix to it's description.
	// It will sort to the beginning and run first
	keys := make([]string, 0)
	for k := range tests {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, desc := range keys {
		t.Run(desc, func(t *testing.T) {
			test := tests[desc]
			result, err := Parse("test.kdl", []byte(test.input))
			assert.NoError(t, err)
			rootNode := mustRootNode(result)
			assert.EqualValues(t, test.expected(), rootNode)
		})
	}
}
