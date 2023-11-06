package kdl

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mustDoc(result interface{}) Document {
	if doc, ok := result.(Document); ok {
		return doc
	}
	if s1, ok := result.([]interface{}); ok {
		if s2, ok := s1[1].([]interface{}); ok {
			return s2[0].(Document)
		}
	}
	panic("Root document not found")
}

func emptyNode() Node {
	return Node{
		Children: []Node{},
		Props:    []Prop{},
		Values:   []Value{},
	}
}

func doc1Node() Document {
	return Document{
		Nodes: []Node{
			{Children: []Node{}, Props: []Prop{}, Values: []Value{}},
		},
	}
}

func TestParse(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected func() Document
	}{
		"just identifier": {
			input: "foo",
			expected: func() Document {
				doc := doc1Node()
				doc.Nodes[0].Identifier = "foo"
				return doc
			},
		},
		"leading newlines": {
			input: "\n\nfoo",
			expected: func() Document {
				doc := doc1Node()
				doc.Nodes[0].Identifier = "foo"
				return doc
			},
		},
		"trailing newlines": {
			input: "foo\n\n",
			expected: func() Document {
				doc := doc1Node()
				doc.Nodes[0].Identifier = "foo"
				return doc
			},
		},
		"newlines before and after": {
			input: "\n\nfoo\n\n",
			expected: func() Document {
				doc := doc1Node()
				doc.Nodes[0].Identifier = "foo"
				return doc
			},
		},
		/*
			"two nodes": {
				input: "alpha\nbravo\n",
				expected: func() Document {
					doc := doc1Node()
					doc.Nodes[0].Identifier = "alpha"
					bravo := emptyNode()
					bravo.Identifier = "bravo"
					doc.Nodes = append(doc.Nodes, bravo)
					return doc
				},
			},
		*/
		"with type": {
			input: "(user)Bill",
			expected: func() Document {
				doc := doc1Node()
				doc.Nodes[0].Identifier = "Bill"
				doc.Nodes[0].Type = "user"
				return doc
			},
		},
		"simple values": {
			input: `foo 0 3.14 "hi" null`,
			expected: func() Document {
				doc := doc1Node()
				doc.Nodes[0].Identifier = "foo"
				doc.Nodes[0].Values = []Value{
					parseValue(0.0),
					parseValue(3.14),
					parseValue("hi"),
					parseValue(nil),
				}
				return doc
			},
		},
		"simple props": {
			input: `foo a=0 b=3.14 c="hi" d=null`,
			expected: func() Document {
				doc := doc1Node()
				doc.Nodes[0].Identifier = "foo"
				doc.Nodes[0].Props = []Prop{
					{Identifier: "a", Value: parseValue(0.0)},
					{Identifier: "b", Value: parseValue(3.14)},
					{Identifier: "c", Value: parseValue("hi")},
					{Identifier: "d", Value: parseValue(nil)},
				}
				return doc
			},
		},
		"slash dash": {
			input: `/-foo a=0 b=3.14 c="hi" d=null`,
			expected: func() Document {
				return Document{Nodes: []Node{}}
			},
		},
	}

	// To debug a particular test case first, add ! as a prefix to it's description.
	// It will sort to the beginning and run first
	keys := make([]string, 0)
	for k := range tests {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, desc := range keys {
		t.Run(desc, func(t *testing.T) {
			test := tests[desc]
			result, err := Parse("test.kdl", []byte(test.input))
			assert.NoError(t, err)
			doc := mustDoc(result)
			assert.EqualValues(t, test.expected(), doc)
		})
	}
}
