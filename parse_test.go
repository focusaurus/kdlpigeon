package kdl

import (
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
	expected := emptyNode()
	expected.Identifier = "foo"
	result, err := Parse("test.kdl", []byte("foo"))
	assert.NoError(t, err)
	rootNode := mustRootNode(result)
	assert.EqualValues(t, expected, rootNode)
}
