package bytetree

import (
	"encoding/binary"
	"testing"
)

var tree *Tree

func init() {
	tree = NewTree()
}

func TestTree(t *testing.T) {
	key0 := []byte("key0")
	key1 := []byte("key1")

	val0 := "value0"
	val1 := "value1"

	//add two values
	tree.GrowLeaf(key0, val0)
	tree.GrowLeaf(key1, val1)

	//lookup value by key
	v, ok := tree.LookupLeaf(key0)
	if !ok || v.(string) != val0 {
		t.Error("failed lookup value")
	}
	t.Log(v)

	v, ok = tree.LookupLeaf(key1)
	if !ok || v.(string) != val1 {
		t.Error("failed lookup value")
	}
	t.Log(v)

	//remove key
	tree.CutLeaf(key0)

	v, ok = tree.LookupLeaf(key0)
	if ok {
		t.Error("failed cut leaf", v)
	}
}

func BenchmarkTreeAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var key = make([]byte, 8)
		binary.PutVarint(key, int64(i))
		tree.GrowLeaf(key, i)
	}
}

func BenchmarkTreeGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var key = make([]byte, 8)
		binary.PutVarint(key, int64(i)%100000)
		tree.LookupLeaf(key)
	}
}
