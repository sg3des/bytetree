package bytetree

type Tree struct {
	root  *branch
	leafs []interface{}
}

func NewTree() *Tree {
	return &Tree{root: &branch{}}
}

func (t *Tree) GrowLeaf(key []byte, v interface{}) {
	t.root.GrowLeaf(key, v)
	t.leafs = append(t.leafs, v)
}

func (t *Tree) LookupLeaf(key []byte) (interface{}, bool) {
	return t.root.LookupLeaf(key)
}

func (t *Tree) PickAllLeafs() []interface{} {
	return t.leafs
}

func (t *Tree) CutLeaf(key []byte) {
	t.root.CutLeaf(key)
}

type branch struct {
	branches [256]*branch
	leaf     interface{}
}

func (b *branch) GrowLeaf(key []byte, v interface{}) {
	if len(key) == 0 {
		b.leaf = v
		return
	}

	if b.branches[key[0]] == nil {
		b.branches[key[0]] = &branch{}
	}

	b.branches[key[0]].GrowLeaf(key[1:], v)
}

func (b *branch) LookupLeaf(key []byte) (interface{}, bool) {
	if len(key) == 0 {
		return b.leaf, b.leaf != nil
	}

	if b.branches[key[0]] == nil {
		return nil, false
	}

	return b.branches[key[0]].LookupLeaf(key[1:])
}

func (b *branch) CutLeaf(key []byte) {
	if len(key) == 0 {
		b.leaf = nil
		return
	}

	if b.branches[key[0]] == nil {
		return
	}

	b.branches[key[0]].CutLeaf(key[1:])
}
