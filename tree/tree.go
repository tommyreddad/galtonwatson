package tree

import "fmt"

// Node represents a node in a rooted ordered tree.
type Node struct {
	Data interface{}

	Parent *Node

	Children []*Node
}

// New creates a new bare node holding any desired sort of data.
func New(data interface{}) *Node {
	return &Node{
		Data: data,
	}
}

// IsRoot returns true if and only if the specified node has no parent.
func (u *Node) IsRoot() bool {
	return u.Parent == nil
}

// IsLeaf returns true and if only if the specified node has no children.
func (u *Node) IsLeaf() bool {
	return u.Children == nil || len(u.Children) == 0
}

// AppendChild creates a parent-child relationship between an input node and
// another given node, appending it to its ordered list of children.
func (u *Node) AppendChild(child *Node) {
	if u.Children == nil {
		u.Children = make([]*Node, 1)
	}
	u.Children = append(u.Children, child)
	child.Parent = u
}

func (u *Node) DFS() {
	if u == nil {
		return
	}
	fmt.Println(u.Data)
	for _, v := range u.Children {
		v.DFS()
	}
}
