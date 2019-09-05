package graph

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	if n := NewNode(1); n == nil {
		t.Error("new node is nil")
	}
}
