# bytetree

lightweight, non-locking, tree, memory, storage.

## logic

Each branch contains up to 256 branches(by byte value) and only one leaf(value). On adding value to tree, byte by byte selects a suitable branch where the value(leaf) will be placed.

## example

```go
//create tree
tree := newTree()

//add value by key to tree
tree.GrowLeaf(key, value)
[...]

//lookup value by key in tree
value, ok := tree.LookupLeaf(key)

//get all values
values := tree.PickAllLeafs()
```






